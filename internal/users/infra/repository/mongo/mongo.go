package mongo

import (
	"context"
	"github.com/guil95/split-entrypoints/internal/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Database
}

func NewCommandMongoRepository(db *mongo.Database) users.UserCommandRepository {
	return repository{db}
}

func (r repository) SaveUser(ctx context.Context, user *users.User) error {
	collection := r.db.Collection("users")

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
