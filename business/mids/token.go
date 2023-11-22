package mids

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shtil/business/validate"
	"shtil/foundation/web"
)

// ApiToken middleware is used for health check endpoint. It looks "ApiToken" request header, and provide api token based authentication.
func ApiToken(token string) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			fmt.Println("TOKEN ", r.Header.Get("Api-Token"))
			fmt.Println("normal token", token)

			if token != r.Header.Get("Api-Token") {
				return validate.RequestError{
					Err:    errors.New("invalid api token"),
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
