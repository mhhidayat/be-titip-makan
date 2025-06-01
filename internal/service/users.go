package service

import (
	"be-titip-makan/domain"
	"be-titip-makan/domain/dto"
	"context"
)

type userService struct {
	usersRepository domain.UsersRepository
}

func NewUsers(userRepository domain.UsersRepository) domain.UserService {
	return &userService{
		usersRepository: userRepository,
	}
}

func (ur userService) Login(ctx context.Context, PhoneNumber string, pasword string) (usersData *dto.UsersData, err error) {
	users, err := ur.usersRepository.FindByPhoneNumberAndPassword(ctx, PhoneNumber, pasword)

	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, nil
	}

	usersData = &dto.UsersData{
		ID:          users.ID,
		Name:        users.Name,
		PhoneNumber: users.PhoneNumber,
	}

	return usersData, nil
}
