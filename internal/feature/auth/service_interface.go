package auth

import (
	"be-titip-makan/internal/feature/user"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, phoneNumber string, password string) (*user.UsersData, error)
}
