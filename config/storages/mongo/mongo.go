package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", os.Getenv("MONGODB_HOST")))
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return client.Database(os.Getenv("DB_DATABASE"))
}
