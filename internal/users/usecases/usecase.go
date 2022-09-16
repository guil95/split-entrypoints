package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/guil95/split-entrypoints/internal/users"
)

type UseCase interface {
	Save(ctx context.Context, user *users.User) error
}

type useCase struct {
	commandRepo users.UserCommandRepository
}

func NewUseCase(repo users.UserCommandRepository) UseCase {
	return useCase{repo}
}

func (uc useCase) Save(ctx context.Context, user *users.User) error {
	user.ID = uuid.New().String()

	return uc.commandRepo.SaveUser(ctx, user)
}
