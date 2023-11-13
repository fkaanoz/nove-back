package models

import "time"

type Order struct {
	ID     string    `db:"id"`
	At     time.Time `db:"at"`
	UserID int       `db:"user_id"`
}
