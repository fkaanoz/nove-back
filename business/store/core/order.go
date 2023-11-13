package core

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shtil/business/store/models"
)

type OrderCore struct {
	DB *sqlx.DB
}

func (oc *OrderCore) ReadLast(limit int) ([]models.Order, error) {
	orders := []models.Order{}

	err := oc.DB.Select(&orders, "SELECT id, at, user_id FROM orders ORDER BY id DESC LIMIT $1 ", limit)
	if err != nil {
		fmt.Println("ERR", err)
		return nil, err
	}

	return orders, nil

}
