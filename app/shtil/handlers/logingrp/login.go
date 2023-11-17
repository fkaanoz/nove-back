package logingrp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"shtil/business/store/core"
	"shtil/business/store/models"
	"shtil/business/validate"
	"shtil/foundation/web"
)

type LoginGrp struct {
	Core core.LoginCore
	Auth *web.Auth
}

func (l *LoginGrp) Register(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	reader := io.LimitReader(r.Body, 1024*1024)
	defer r.Body.Close()

	//body, err := io.ReadAll(b)
	//if err != nil {
	//	return errors.New("body read error")
	//}

	u := models.User{}

	err := web.DecodeJSONBody(reader, &u, r)
	if err != nil {
		fmt.Println("decode err", err)
		return err
	}

	err = l.Core.SaveUser(&u)
	if err != nil {
		fmt.Println("save err", err)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it is ok"))

	return nil
}

func (l *LoginGrp) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// decode body
	c := struct {
		Email    string
		Password string
	}{}

	reader := io.LimitReader(r.Body, 1024*1024)
	defer r.Body.Close()

	err := web.DecodeJSONBody(reader, &c, r)
	if err != nil {
		return errors.New("decode error")
	}

	// compare hashes
	ok := l.Core.RetrieveAndComparePassword(c.Email, c.Password)
	if !ok {
		return validate.RequestError{
			Err:    errors.New("invalid credentials"),
			Fields: nil,
		}
	}

	// create token
	token, err := l.Auth.GenerateToken(24)
	if err != nil {
		return errors.New("create token error")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`bearer : %s`, token)))

	return nil
}
