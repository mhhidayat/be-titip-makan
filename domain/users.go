package domain

import (
	"be-titip-makan/domain/dto"
	"context"
	"database/sql"
)

type Users struct {
	ID          string       `db:"id"`
	Name        string       `db:"name"`
	PhoneNumber string       `db:"phone_number"`
	Password    string       `db:"password"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}

type UsersRepository interface {
	FindByPhoneNumberAndPassword(ctx context.Context, phoneNumber string, password string) (*Users, error)
}

type UserService interface {
	Login(ctx context.Context, phoneNumber string, password string) (*dto.UsersData, error)
}
