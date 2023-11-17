package core

import (
	"context"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"shtil/business/store/models"
	"time"
)

type LoginCore struct {
	DB *sqlx.DB
}

func (l *LoginCore) SaveUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	_, err = l.DB.NamedExecContext(ctx, `INSERT INTO users(name, surname, email, password) VALUES(:name, :surname, :email, :password)`, user)
	if err != nil {
		return err
	}

	return nil
}

func (l *LoginCore) RetrieveAndComparePassword(email string, password string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	u := models.User{}

	err := l.DB.GetContext(ctx, &u, `SELECT password FROM users WHERE email=$1`, email)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}

	return true
}
