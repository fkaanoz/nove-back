package usergrp

import (
	"context"
	"errors"
	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
	"io"
	"net/http"
	credis "shtil/app/redis"
	"shtil/business/store/core"
	"shtil/business/validate"
	"shtil/foundation/web"
	"strconv"
)

type UserGrp struct {
	Redis  *credis.Redis
	Logger *zap.SugaredLogger
	Core   *core.UserCore
	Auth   *web.Auth
}

func (ug *UserGrp) UserByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// check token and return for test
	//
	//valid := ug.Auth.ValidateToken(r.Header.Get("token"))
	//fmt.Print("token validation ")

	urlParams := httptreemux.ContextData(r.Context()).Params()
	userID, ok := urlParams["id"]
	if !ok {
		return validate.RequestError{
			Err:    errors.New("id is not provided"),
			Fields: []string{"id"},
		}
	}

	id, err := strconv.Atoi(userID)
	if err != nil {
		return validate.RequestError{
			Err:    errors.New("not proper id"),
			Fields: []string{"id"},
		}
	}

	user, err := ug.Core.ReadByID(id)
	if err != nil {
		return validate.RequestError{
			Err: errors.New("user cannot be found"),
		}
	}

	web.Respond(w, user, http.StatusOK)
	return nil
}

func (ug *UserGrp) UserByName(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	reader := io.LimitReader(r.Body, 1024)
	defer r.Body.Close()

	b := struct {
		Name string `json:"name"`
	}{}

	err := web.DecodeJSONBody(reader, &b, r)
	if err != nil {
		return err
	}

	user, err := ug.Core.ReadByName(b.Name)
	if err != nil {
		return err
	}

	web.Respond(w, user, http.StatusOK)
	return nil
}

func (ug *UserGrp) UserByEmail(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return nil
}
