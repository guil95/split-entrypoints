package commands

import (
	"github.com/guil95/split-entrypoints/config/entrypoints/kafka"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewConsumersCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "consumers",
		Short: "Start consumers",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go kafka.RunKafkaConsumers(quit)

			err := <-quit

			zap.S().Fatal("Error grpc server:", err)
		},
	}
}
