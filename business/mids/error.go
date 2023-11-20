package mids

import (
	"context"
	"net/http"
	"shtil/business/validate"
	"shtil/foundation/web"
)

func Error() web.Middleware {

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// classify errors and handle them separately.
			if err := handler(ctx, w, r); err != nil {
				var er validate.ErrorResponse
				var status int

				switch act := validate.Cause(err).(type) {
				case validate.RequestError:
					status = http.StatusBadRequest
					er = validate.ErrorResponse{
						Error:  act.Error(),
						Fields: act.Fields,
						Reason: "bad request",
					}
				case validate.NotFoundError:
					status = http.StatusNotFound
					er = validate.ErrorResponse{
						Error:  act.Message,
						Fields: nil,
						Reason: "not found",
					}
				default:
					status = http.StatusInternalServerError
					er = validate.ErrorResponse{
						Error:  act.Error(),
						Fields: nil,
						Reason: "internal service error",
					}
				}

				if err := web.Respond(w, er, status); err != nil {
					return err
				}

				return nil
			}
			return nil
		}
		return h
	}
	return m
}
