package domain

import "context"

type UserRepository interface {
	All(ctx context.Context) ([]User, error)
	Find(ctx context.Context, in User) ([]User, error)
	FindById(ctx context.Context, in string) (*User, error)
	Create(ctx context.Context, in User) (string, error)
	Update(ctx context.Context, in User) error
	Delete(ctx context.Context, in string) error
}
