package transactiongrp

import (
	"context"
	"errors"
	"github.com/dimfeld/httptreemux/v5"
	"net/http"
	"shtil/business/store/core"
	"shtil/business/validate"
	"shtil/foundation/web"
	"strconv"
)

type TransactionGrp struct {
	Core *core.TransactionCore
}

func (tg *TransactionGrp) GetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	urlParams := httptreemux.ContextData(r.Context()).Params()
	txID, ok := urlParams["id"]
	if !ok {
		return validate.RequestError{
			Err:    errors.New("tx id is not provided"),
			Fields: nil,
		}
	}

	id, err := strconv.Atoi(txID)
	if err != nil {
		return validate.RequestError{
			Err:    errors.New("improper id is provided"),
			Fields: nil,
		}
	}

	tx, err := tg.Core.ReadByID(id)
	if err != nil {
		return err
	}

	web.Respond(w, tx, http.StatusOK)
	return nil
}
