package models

type Transaction struct {
	ID int `db:"id"`
	Order
	User
}
