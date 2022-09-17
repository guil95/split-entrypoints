package main

import (
	"context"
	"github.com/guil95/outbox"
	"github.com/guil95/split-entrypoints/cmd/workers"
	"github.com/guil95/split-entrypoints/config/brokers/kafka"
	"github.com/guil95/split-entrypoints/config/storages/mongo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	o := outbox.NewOutbox(
		outbox.NewMongoStorage(mongo.Connect()),
		outbox.NewKafkaProducer(kafka.RetrieveProducer()),
	)
	go o.Listen(ctx)

	rootCmd := &cobra.Command{
		Use: "split-entry-points",
	}

	rootCmd.AddCommand(workers.NewGrpcWorker())
	rootCmd.AddCommand(workers.NewHTTPWorker())
	rootCmd.AddCommand(workers.NewConsumersWorker())
	_ = rootCmd.Execute()
}
