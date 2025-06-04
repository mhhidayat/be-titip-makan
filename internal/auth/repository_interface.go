package auth

import (
	"be-titip-makan/internal/user"
	"context"
)

type AuthRepository interface {
	Login(ctx context.Context, phoneNumber string, password string) (*user.Model, error)
}
