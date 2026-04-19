package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
			style := "--bin"
			if libFlag {
				style = "--lib"
			}

			fmt.Println("Preparing your rust project...")

			execCmd := exec.Command(
				"cargo",
				"init",
				projectName,
				style,
			)

			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr

			if err := execCmd.Run(); err != nil {
				return fmt.Errorf("failed to run cargo init for %q: %w", projectName, err)
			}

			if err := runFazInit(projectName); err != nil {
				return err
			}

			if err := appendGitInfoExclude(projectName); err != nil {
				return err
			}

			src := "assets/justfile_rust"
			out := filepath.Join(projectName, "justfile")

			if err := copyEmbFile(assets, src, out); err != nil {
				return err
			}

			fmt.Println("\nDone!!")

			return nil
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
