package user

import "database/sql"

type Model struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	PhoneNumber string       `db:"phone_number"`
	Password    string       `db:"password"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
