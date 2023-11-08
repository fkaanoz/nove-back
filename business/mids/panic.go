package mids

import (
	"context"
	"errors"
	"net/http"
	"shtil/foundation/web"
)

func Panic() web.Middleware {

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = errors.New("panicking error")
				}
			}()
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
