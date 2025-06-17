package auth

import (
	"be-titip-makan/internal/feature/user"
	"context"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) (*user.UsersData, error)
}
