package main

import (
	"github.com/guil95/split-entrypoints/cmd/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "split-entry-points"}

	rootCmd.AddCommand(commands.NewGrpcCommand())
	rootCmd.AddCommand(commands.NewHTTPCommand())
	rootCmd.AddCommand(commands.NewConsumersCommand())
	_ = rootCmd.Execute()
}
