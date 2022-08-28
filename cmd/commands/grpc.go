package commands

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/guil95/split-entrypoints/config/entrypoints/grpc"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewGrpcCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "grpc",
		Short: "Start grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go grpc.RunGrpcServer(quit)

			err := <-quit

			zap.S().Fatal("Error grpc server:", err)
		},
	}
}
