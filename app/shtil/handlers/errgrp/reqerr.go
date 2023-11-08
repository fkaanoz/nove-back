package errgrp

import (
	"context"
	"errors"
	"net/http"
	"shtil/business/validate"
)

func ReqErrSimulation(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return validate.RequestError{
		Err: errors.New("request err msg"),
	}
}
