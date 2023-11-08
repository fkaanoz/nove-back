package mids

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"shtil/foundation/web"
)

func Logger(log *zap.SugaredLogger) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			log.Infow("REQUEST", "status", "started", "traceID", web.GetTraceID(ctx))

			if err := handler(ctx, w, r); err != nil {
				log.Errorw("REQUEST", "status", "error", "ERROR", err, "traceID", web.GetTraceID(ctx))
				return err
			}

			log.Errorw("REQUEST", "status", "successful", "traceID", web.GetTraceID(ctx))
			return nil
		}

		return h
	}

	return m
}
