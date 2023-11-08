package validate

import "errors"

type ErrorResponse struct {
	Error  string   `json:"error"`
	Fields []string `json:"fields"`
	Reason string   `json:"reason"`
}

type RequestError struct {
	Err    error
	Fields []string
}

func (re RequestError) Error() string {
	return re.Err.Error()
}

func NewRequestError(err error, fields []string) *RequestError {
	return &RequestError{
		Err:    err,
		Fields: fields,
	}
}

// Cause finds the root of the provided error.
func Cause(err error) error {
	root := err
	for {
		if err := errors.Unwrap(err); err == nil {
			return root
		}
		root = err
	}
}
