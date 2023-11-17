package mids

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"shtil/business/validate"
	"shtil/foundation/web"
)

func Auth(auth *web.Auth) web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// handle all auth logic here and call the "handler"
			// TODO : PARSE and Validate

			token := r.Header.Get("token")

			fmt.Println("provided token by client", token)

			if !auth.ValidateToken(token) {
				fmt.Println("non valid token in MID")
				return validate.RequestError{
					Err:    errors.New("token is not provided"),
					Fields: nil,
				}
			}

			fmt.Println("auth mid is called")

			if err := handler(ctx, w, r); err != nil {
				return err
			}

			return nil
		}
		return h
	}
	return m
}
