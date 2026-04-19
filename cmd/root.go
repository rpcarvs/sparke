// Package cmd holds the CLI base command
package cmd

import (
	"embed"
	"fmt"

	"github.com/spf13/cobra"
)

//go:embed assets/*
var assets embed.FS

// NewRootCmd builds the sparke CLI command tree.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
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

	addGoCommand(rootCmd)
	addRustCommand(rootCmd)
	addPythonCommand(rootCmd)

	return rootCmd
}
