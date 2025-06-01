package repository

import (
	"be-titip-makan/domain/user"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type usersRepository struct {
	db *goqu.Database
}

func NewUsers(con *sql.DB) user.UsersRepository {
	return &usersRepository{
		db: goqu.New("default", con),
	}
}

func (ur usersRepository) FindByPhoneNumberAndPassword(ctx context.Context, phoneNumber string, password string) (*user.Model, error) {
	result := user.Model{}

	expresion := goqu.Ex{
		"phone_number": phoneNumber,
		"password":     password,
	}

	dataset := ur.db.From("users").
		Where(expresion).
		Limit(1)

	found, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return &result, nil
}
