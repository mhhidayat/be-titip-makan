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

func (ur authService) Login(ctx context.Context, username string, pasword string) (usersData *user.UsersData, err error) {
	users, err := ur.usersRepository.Login(ctx, username, pasword)

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, nil
	}

	usersData = &user.UsersData{
		ID:          users.ID,
		Username:    users.Username,
		Name:        users.Name,
		PhoneNumber: users.PhoneNumber,
	}

	return usersData, nil
}
