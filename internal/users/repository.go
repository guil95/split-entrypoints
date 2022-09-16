package users

import (
	"context"
)

type UserCommandRepositoryError struct{}

func (ce UserCommandRepositoryError) Error() string {
	return "Error to save user"
}

type UserCommandRepository interface {
	SaveUser(ctx context.Context, user *User) error
}

type UserQueryRepositoryError struct{}

func (ce UserQueryRepositoryError) Error() string {
	return "Error to find user"
}

type UserQueryRepository interface {
	FindUserByName(ctx context.Context, name string) error
	FindUserByID(ctx context.Context, id string) error
}
