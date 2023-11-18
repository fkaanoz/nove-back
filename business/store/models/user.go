package models

type User struct {
	ID       string `db:"id" json:"-"`
	Name     string `db:"name" json:"name"`
	Surname  string `db:"surname" json:"surname"`
	Password string `db:"password" json:"-"`
	Email    string `db:"email" json:"email"`
}
