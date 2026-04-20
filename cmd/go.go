package cmd

import (
	"github.com/spf13/cobra"
)

func newGoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "go [project_name]",
		Short: "Scaffold a go project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			return scaffoldGoProject(projectName, projectName)
		},
	}
}

// addGoCommand wires the Go scaffolding command into the root tree.
func addGoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newGoCmd())
}
