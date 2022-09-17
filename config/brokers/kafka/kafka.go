package kafka

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func RetrieveProducer() *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        os.Getenv("KAFKA_HOST"),
		"delivery.timeout.ms":      600000,
		"linger.ms":                10000,
		"message.send.max.retries": 10000000,
		"batch.num.messages":       1,
		"enable.idempotence":       true,
	})

	if err != nil {
		panic(err)
	}

	return producer
}
