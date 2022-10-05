package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/guil95/outbox"
	"github.com/guil95/split-entrypoints/internal/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

func NewCommandMongoRepository(db *mongo.Client) users.UserCommandRepository {
	return repository{db}
}

func (r repository) SaveUser(ctx context.Context, user *users.User) error {
	session, err := r.db.StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(ctx)

	transactional := func(sessCtx mongo.SessionContext) (interface{}, error) {
		db := r.db.Database(os.Getenv("DB_DATABASE"))
		collectionUsers := db.Collection("users")
		collectionOutbox := db.Collection("outbox")

		_, err := collectionUsers.InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}

		message, err := json.Marshal(user)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return nil, err
		}

		_, err = collectionOutbox.InsertOne(ctx, &outbox.Model{
			IdempotencyID: user.ID,
			Message:       string(message),
			Topic:         "user_created",
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, transactional)
	if err != nil {
		return err
	}

	return nil
}
