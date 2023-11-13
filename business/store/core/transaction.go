package core

import (
	"github.com/jmoiron/sqlx"
	"shtil/business/store/models"
)

type TransactionCore struct {
	DB *sqlx.DB
}

func (t *TransactionCore) ReadByID(id int) (*models.Transaction, error) {
	tx := models.Transaction{}

	err := t.DB.Get(&tx, "SELECT transactions.id, users.name, users.surname, orders.at FROM transactions LEFT JOIN users ON transactions.id = users.id LEFT JOIN orders ON transactions.id = orders.id  WHERE transactions.id=$1", id)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}
