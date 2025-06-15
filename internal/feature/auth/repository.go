package auth

import (
	"be-titip-makan/internal/feature/user"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type authRepository struct {
	db *goqu.Database
}

func NewAuthRepository(con *sql.DB) AuthRepository {
	return &authRepository{
		db: goqu.New("default", con),
	}
}

func (ur authRepository) Login(ctx context.Context, phoneNumber string, password string) (*user.Model, error) {
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
