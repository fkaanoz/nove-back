package ordergrp

import (
	"context"
	"go.uber.org/zap"
	"net/http"
	"shtil/business/store/core"
	"shtil/business/validate"
	"shtil/foundation/web"
)

type OrderGrp struct {
	Logger *zap.SugaredLogger
	Core   *core.OrderCore
}

func (og *OrderGrp) Last20Orders(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	orders, err := og.Core.ReadLast(20)
	if err != nil {
		return validate.RequestError{
			Err:    err,
			Fields: nil,
		}
	}

	web.Respond(w, orders, http.StatusOK)

	return nil
}
