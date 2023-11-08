package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	ckafka "shtil/app/kafka"
	credis "shtil/app/redis"
	"shtil/app/shtil/debug"
	"shtil/app/shtil/handlers"
	bkafka "shtil/business/kafka"
	"shtil/business/logger"
	"shtil/business/store"
	"shtil/foundation/web"
	"syscall"
	"time"
)

var Service = "sthil"
var ServiceVersion = "v1.0.0"
var BUILD = "dev"

func main() {
	zlog, err := logger.InitLogger(Service)
	if err != nil {
		log.Fatal("init log err : ", err)
	}

	if err := run(zlog); err != nil {
		zlog.Errorw("RUN_ERROR", "ERROR", err)
	}
}

func run(log *zap.SugaredLogger) error {

	// init config
	cfg := struct {
		conf.Version
		Web struct {
			Addr            string        `conf:"default:0.0.0.0:4001"`
			ReadTimeout     time.Duration `conf:"default:20s"`
			WriteTimeout    time.Duration `conf:"default:20s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
		Debug struct {
			Addr string `conf:"default:0.0.0.0:4002"`
		}
		Kafka struct {
			BootstrapServers string `conf:"default:0.0.0.0:9092,0.0.0.0:9093"`
			GroupID          string `conf:"default:test-group"`
			AutoCommit       bool   `conf:"default:false"`
			AutoOffsetReset  string `conf:"default:earliest"`
		}
		Socket struct{}
		Redis  struct {
			Addr     string `conf:"default:0.0.0.0:6379"`
			Password string `conf:"default:''"`
			DB       int    `conf:"default:0"`
		}
		DB struct {
			Username    string `conf:"default:postgres"`
			Password    string `conf:"default:fkaanoz"`
			Host        string `conf:"default:localhost"`
			Database    string `conf:"default:test_db"`
			MaxIdleConn int    `conf:"default:25"`
			MaxOpenConn int    `conf:"default:25"`
		}
	}{
		Version: conf.Version{
			Build: BUILD,
			Desc:  "sthil for test",
		},
	}

	prefix := "STHIL"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return err
	}

	serverCh := make(chan error, 1)
	shutdownCh := make(chan os.Signal)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	// redis
	redisOptions := &redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	}
	redisClient := credis.NewRedisClientWithOptions(redisOptions)

	// database
	dbConfig := store.DBConfig{
		Username:    cfg.DB.Username,
		Password:    cfg.DB.Password,
		Host:        cfg.DB.Host,
		Database:    cfg.DB.Database,
		SSLRequire:  false,
		Timezone:    "Europe/Istanbul",
		MaxIdleConn: cfg.DB.MaxIdleConn,
		MaxOpenConn: cfg.DB.MaxOpenConn,
	}

	conn, err := store.Connect(dbConfig)
	if err != nil {
		return err
	}

	api := http.Server{
		Addr: cfg.Web.Addr,
		Handler: handlers.NewApp(&web.AppConfig{
			Logger:      log,
			Redis:       redisClient,
			DB:          conn,
			ServerErrCh: serverCh,
		}),
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
	}

	// debug
	go func() {
		log.Infow("DEBUG", "status", "started", "port", cfg.Debug.Addr)
		defer log.Infow("DEBUG", "status", "stopped", "port", cfg.Debug.Addr)
		if err := http.ListenAndServe(cfg.Debug.Addr, debug.DebugApi()); err != nil {
			fmt.Println("debug api err : ", err)
		}
	}()

	// api
	go func() {
		log.Infow("API", "status", "started", "port", cfg.Web.Addr)
		serverCh <- api.ListenAndServe()
	}()

	// kafka
	consumer, err := ckafka.NewKafkaConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  cfg.Kafka.BootstrapServers,
		"group.id":           cfg.Kafka.GroupID,
		"enable.auto.commit": cfg.Kafka.AutoCommit,
		"auto.offset.reset":  cfg.Kafka.AutoOffsetReset,
	}, log, []string{bkafka.TestTopic, bkafka.AnotherTopic})

	//  TODO : TOPIC LIST SHOULD NOT BE A SLICE. CREATE STRUCTS PER TOPICS AND HANDLERS.

	go func() {
		log.Infow("KAFKA", "status", "listening")
		defer log.Errorw("KAFKA", "status", "stopped")
		consumer.Run()
	}()

	select {
	case <-shutdownCh:
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		log.Infow("SHUTDOWN", "status", "started")
		if err := api.Shutdown(ctx); err != nil {
			log.Errorw("SHUTDOWN", "status", "gracefully shutdown is not possible. it is forced to stop.")
			api.Close()
			return nil
		}
		log.Infow("SHUTDOWN", "status", "finished")
	case err := <-serverCh:
		log.Errorw("API", "ERROR", err)
	}

	return nil
}

func subscribe() []string {
	return []string{"test-topic"}
}
