package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brandoyts/go-clean/internal/domain"
	"github.com/brandoyts/go-clean/internal/service"
	"github.com/brandoyts/go-clean/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testUserService struct {
	description string
	input       any
	expected    any
	err         error
}

func TestFindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	userService := service.NewUserService(mockUserRepo)

	mockUserRepo.EXPECT().FindById(context.Background(), "000-111").Return(&domain.User{
		ID: "000-111",
	}, nil)

	result, err := userService.GetUserById(context.Background(), "000-111")
	assert.NoError(t, err)

	assert.Equal(t, "000-111", result.ID)
}

func TestFinAllUsers(t *testing.T) {
	tests := []testUserService{
		{
			description: "should return 3 users",
			input: []domain.User{
				{ID: "000-111"},
				{ID: "000-222"},
				{ID: "000-333"},
			},
			expected: 3,
			err:      nil,
		},
		{
			description: "should return 1 user",
			input: []domain.User{
				{ID: "000-444"},
			},
			expected: 1,
			err:      nil,
		},
		{
			description: "should return error",
			input:       nil,
			expected:    0,
			err:         errors.New("custom error"),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			userService := service.NewUserService(mockUserRepo)

			mockUserRepo.EXPECT().All(context.Background()).Return(test.input, test.err)
			got, err := userService.GetAllUser(context.Background())

			assert.Exactly(t, test.err, err)
			assert.Equal(t, test.expected, len(got))
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []testUserService{
		{
			description: "should create a user with inserted id of 000-000",
			input: domain.User{
				Email: "user1@mail.com",
			},
			expected: "000-000",
			err:      nil,
		},
		{
			description: "should create a user with inserted id of 111-111",
			input: domain.User{
				Email: "user2@mail.com",
			},
			expected: "111-112",
			err:      nil,
		},
		{
			description: "should return error",
			input:       domain.User{},
			expected:    "",
			err:         errors.New("custom error"),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			userService := service.NewUserService(mockUserRepo)

			mockUserRepo.EXPECT().Create(context.Background(), test.input).Return(test.expected, test.err)

			got, err := userService.CreateUser(context.Background(), test.input.(domain.User))

			assert.Exactly(t, test.err, err)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []testUserService{
		{
			description: "should be able to update a user without error",
			input:       "999090-000",
			err:         nil,
		},
		{
			description: "should return error",
			input:       "999090-000",
			err:         errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			userService := service.NewUserService(mockUserRepo)

			mockUserRepo.EXPECT().Delete(context.Background(), test.input).Return(test.err)

			err := userService.DeleteUser(context.Background(), test.input.(string))
			assert.Exactly(t, test.err, err)
		})
	}
}
