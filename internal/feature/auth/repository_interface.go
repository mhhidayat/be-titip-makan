package auth

import (
	"be-titip-makan/internal/feature/user"
	"context"
)

type AuthRepository interface {
	Login(ctx context.Context, username string, password string) (*user.Model, error)
}
