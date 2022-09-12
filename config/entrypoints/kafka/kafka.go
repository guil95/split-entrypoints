package kafka

import (
	"github.com/guil95/split-entrypoints/internal/users/infra/broker/kafka"
	"os"
)

func RunKafkaConsumers(quit chan os.Signal) {
	// config kafka

	kafka.Listen(quit)
}
