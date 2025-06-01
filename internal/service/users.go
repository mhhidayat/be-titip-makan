package service

import (
	"be-titip-makan/domain/user"
	"context"
)

type userService struct {
	usersRepository user.UsersRepository
}

func NewUsers(userRepository user.UsersRepository) user.UserService {
	return &userService{
		usersRepository: userRepository,
	}
}

func (ur userService) Login(ctx context.Context, PhoneNumber string, pasword string) (usersData *user.UsersData, err error) {
	users, err := ur.usersRepository.FindByPhoneNumberAndPassword(ctx, PhoneNumber, pasword)

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
