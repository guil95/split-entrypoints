package main

import (
	"github.com/guil95/split-entrypoints/cmd/workers"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "split-entry-points"}

	rootCmd.AddCommand(workers.NewGrpcWorker())
	rootCmd.AddCommand(workers.NewHTTPWorker())
	rootCmd.AddCommand(workers.NewConsumersWorker())
	_ = rootCmd.Execute()
}
