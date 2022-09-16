package command

import (
	"context"
	"go.uber.org/zap"

	"github.com/guil95/split-entrypoints/internal/users"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewMysqlCommandRepository(db *sqlx.DB) users.UserCommandRepository {
	return repository{db}
}

func (r repository) SaveUser(ctx context.Context, user *users.User) error {
	_, err := r.db.NamedExecContext(ctx, "INSERT INTO users (name, age) VALUES (:name, :age)", user)

	if err != nil {
		zap.S().Errorf("error repository: %v", err)
		return users.UserCommandRepositoryError{}
	}

	return nil
}
