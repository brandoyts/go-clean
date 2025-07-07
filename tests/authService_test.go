package tests

import (
	"context"
	"testing"

	"github.com/brandoyts/go-clean/internal/domain"
	"github.com/brandoyts/go-clean/internal/service"
	"github.com/brandoyts/go-clean/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testAuthService struct {
	description string
	input       any
	expected    any
	err         error
}

func TestRegisterUser(t *testing.T) {
	tests := []testAuthService{
		{
			description: "should register a user with email user1@mail.com",
			input:       domain.User{Email: "user1@mail.com", Password: "password1"},
			expected:    "000-000",
			err:         nil,
		},
		{
			description: "should return error on no password provided",
			input:       domain.User{Email: "user1@mail.com"},
			expected:    "",
			err:         domain.ErrUserHashPassword,
		},
		{
			description: "should return error on empty string password",
			input:       domain.User{Email: "user1@mail.com", Password: ""},
			expected:    "",
			err:         domain.ErrUserHashPassword,
		},
		{
			description: "should return error on password with only whitespaces",
			input:       domain.User{Email: "user1@mail.com", Password: "  "},
			expected:    "",
			err:         domain.ErrUserHashPassword,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			authService := service.NewAuthService(mockUserRepo)

			if test.err == nil {
				mockUserRepo.
					EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(test.expected, test.err)
			}

			err := authService.Register(context.Background(), test.input.(domain.User))

			assert.ErrorIs(t, err, test.err)
		})
	}

}
