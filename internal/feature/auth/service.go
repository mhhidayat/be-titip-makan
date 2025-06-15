package auth

import (
	"be-titip-makan/internal/feature/user"
	"context"
)

type authService struct {
	usersRepository AuthRepository
}

func NewAuthService(userRepository AuthRepository) AuthService {
	return &authService{
		usersRepository: userRepository,
	}
}

func (ur authService) Login(ctx context.Context, PhoneNumber string, pasword string) (usersData *user.UsersData, err error) {
	users, err := ur.usersRepository.Login(ctx, PhoneNumber, pasword)

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, nil
	}

	usersData = &user.UsersData{
		ID:          users.ID,
		Name:        users.Name,
		PhoneNumber: users.PhoneNumber,
	}

	return usersData, nil
}
