package web

import (
	"context"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"net/http"
	credis "shtil/app/redis"
)

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

// AppConfig should be implemented to reduce parameter count of NewApp constructor.
type AppConfig struct {
	Logger      *zap.SugaredLogger
	Redis       *credis.Redis
	DB          *sqlx.DB
	ServerErrCh chan error
}

type App struct {
	*httptreemux.ContextMux
	Logger      *zap.SugaredLogger
	ServerErrCh chan error
	Middlewares []Middleware
	Redis       *credis.Redis
	DB          *sqlx.DB
}

func (a *App) Handle(method string, path string, handler Handler, middlewares ...Middleware) {

	// wrap with application wise middlewares
	handler = wrapMiddlewares(handler, a.Middlewares...)

	// wrap with handler specific middlewares
	handler = wrapMiddlewares(handler, middlewares...)

	h := func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(context.Background(), CtxKey, Values{
			TraceID: uuid.New().String(),
		})

		if err := handler(ctx, w, r); err != nil {
			a.ServerErrCh <- err
		}
	}

	a.ContextMux.Handle(method, path, h)
}
