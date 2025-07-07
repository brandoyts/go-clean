package service

import (
	"context"
	"log"

	"github.com/brandoyts/go-clean/internal/domain"
)

type AuthService struct {
	userRepository domain.UserRepository
}

func NewAuthService(userRepository domain.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (as *AuthService) Register(ctx context.Context, user domain.User) error {
	err := user.HashPassword()
	if err != nil {
		return err
	}

	result, err := as.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}
