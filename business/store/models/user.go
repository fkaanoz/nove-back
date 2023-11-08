package models

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Surname  string `db:"surname"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
