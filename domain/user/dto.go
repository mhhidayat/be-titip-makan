package user

import (
	"context"
	"database/sql"
)

type Model struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	PhoneNumber string       `db:"phone_number"`
	Password    string       `db:"password"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type UsersData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UsersRepository interface {
	FindByPhoneNumberAndPassword(ctx context.Context, phoneNumber string, password string) (*Model, error)
}

type UserService interface {
	Login(ctx context.Context, phoneNumber string, password string) (*UsersData, error)
}
