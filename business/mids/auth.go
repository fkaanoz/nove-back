package mids

import (
	"context"
	"net/http"
	"shtil/business/validate"
	"shtil/foundation/web"
)

func Auth(auth *web.Auth) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			token := r.Header.Get("token")

			if err := auth.ValidateToken(token); err != nil {
				return validate.RequestError{
					Err:    err,
					Fields: nil,
				}
			}

			if err := handler(ctx, w, r); err != nil {
				return err
			}

			return nil
		}
		return h
	}
	return m
}
