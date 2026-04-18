// Package cmd holds the CLI base command
package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed assets/*
var assets embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sparke",
	Short: "A simple CLI to scaffold Rust, Go and Python projects",
	Long: `The CLI will set a minimal dir structure and copy a
	corresponding justfile with common recipes for the selected
	language.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// prints the help message by default
		err := cmd.Help()
		if err != nil {
			return fmt.Errorf("something went wrong %w", err)
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
