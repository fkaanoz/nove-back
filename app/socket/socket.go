package socket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var Upgrader websocket.Upgrader

type Server struct {
	Addr         string
	Logger       *zap.SugaredLogger
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) Run() error {
	g := gin.New()

	h := http.Server{
		Addr:         s.Addr,
		Handler:      v1(g),
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	return h.ListenAndServe()
}

func v1(engine *gin.Engine) *gin.Engine {
	v := engine.Group("/ws/v1")
	v.GET("/connect", Connect)

	return engine
}
