package core

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"shtil/business/store/models"
)

type UserCore struct {
	DB *sqlx.DB
}

func (u *UserCore) ReadByID(id int) (*models.User, error) {
	user := &models.User{}

	err := u.DB.Get(user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserCore) ReadByName(name string) (*models.User, error) {
	user := &models.User{}

	err := u.DB.Get(user, "SELECT * FROM users WHERE name=$1", name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
