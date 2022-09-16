package workers

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/guil95/split-entrypoints/config/entrypoints/http"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewHTTPWorker() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Start http server",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			go http.RunHTTPServer(quit)

			err := <-quit

			zap.S().Fatal("Error http server:", err)
		},
	}
}
