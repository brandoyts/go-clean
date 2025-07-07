package service

import (
	"context"

	"github.com/brandoyts/go-clean/internal/domain"
)

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) GetAllUser(ctx context.Context) ([]domain.User, error) {
	result, err := u.userRepository.All(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserService) GetUserById(ctx context.Context, id string) (*domain.User, error) {

	result, err := u.userRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *UserService) CreateUser(ctx context.Context, user domain.User) (string, error) {

	result, err := u.userRepository.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id string) error {

	err := u.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
