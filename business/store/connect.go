package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/url"
)

type DBConfig struct {
	Username    string
	Password    string
	Host        string
	Database    string
	SSLRequire  bool
	Timezone    string
	MaxIdleConn int
	MaxOpenConn int
}

func Connect(config DBConfig) (*sqlx.DB, error) {
	sslmode := "require"
	if !config.SSLRequire {
		sslmode = "disable"
	}

	// example connection string ->  postgresql://username:password@198.51.100.22:3333/sales?connect_timeout=10&sslmode=require&target_session_attrs=primary
	connStr := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.Username, config.Password),
		Host:     config.Host,
		Path:     config.Database,
		RawQuery: fmt.Sprintf("sslmode=%s&timezone=%s", sslmode, config.Timezone),
	}

	conn, err := sqlx.Connect("postgres", connStr.String())
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(config.MaxIdleConn)
	conn.SetMaxOpenConns(config.MaxOpenConn)

	return conn, nil
}
