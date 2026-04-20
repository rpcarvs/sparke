package cmd

import (
	"github.com/spf13/cobra"
)

func newRustCmd() *cobra.Command {
	var libFlag bool

	rustCmd := &cobra.Command{
		Use:   "rust [project_name]",
		Short: "Scaffold a rust project using cargo",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			projectType := rustBinaryCrate
			if libFlag {
				projectType = rustLibraryCrate
			}

			return scaffoldRustProject(projectName, projectType)
		},
	}

	rustCmd.Flags().BoolVar(&libFlag, "lib", false, "Create a lib crate")
	rustCmd.Flags().Bool("bin", false, "Create a binary crate (default)")

	return rustCmd
}

// addRustCommand wires the Rust scaffolding command into the root tree.
func addRustCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newRustCmd())
}
