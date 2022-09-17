package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/guil95/outbox"
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

	message, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	collectionOutbox := r.db.Collection("outbox")

	_, err = collectionOutbox.InsertOne(ctx, &outbox.Model{
		IdempotencyID: user.ID,
		Message:       string(message),
		Topic:         "users",
		Event:         "user_saved",
	})
	if err != nil {
		return err
	}

	return nil
}
