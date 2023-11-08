package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shtil/business/validate"
	"strings"
)

// DecodeJSONBody is helper for decoding request body.
func DecodeJSONBody(reader io.Reader, dest interface{}, r *http.Request) error {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		return validate.RequestError{
			Err: errors.New("improper content-type"),
		}
	}

	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields() // It makes decoder to something can return error when additional fields exists in the json.

	err := decoder.Decode(dest)
	if err != nil {
		var jsonSyntaxErr *json.SyntaxError
		var unmarshallTypeErr *json.UnmarshalTypeError

		switch {
		case errors.Is(err, io.ErrUnexpectedEOF):
			return validate.RequestError{
				Err:    errors.New("unexpected EOF"),
				Fields: nil,
			}
		case errors.Is(err, io.EOF):
			return validate.RequestError{
				Err:    errors.New("normal EOF"),
				Fields: nil,
			}
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			return validate.RequestError{
				Err:    errors.New("unexpected fields"),
				Fields: nil,
			}
		case errors.As(err, &jsonSyntaxErr):
			return validate.RequestError{
				Err:    errors.New("syntax error"),
				Fields: nil,
			}
		case errors.As(err, &unmarshallTypeErr):
			return validate.RequestError{
				Err:    errors.New("unmarshall error"),
				Fields: nil,
			}
		default:
			return errors.New("unknown err")
		}
	}

	return nil
}
