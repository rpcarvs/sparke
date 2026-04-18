package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rsLibFlag bool

var rustCmd = &cobra.Command{
	Use:   "rust [project_name]",
	Short: "Scaffold a rust project using cargo",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		// default
		style := "--bin"
		if rsLibFlag {
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

		// executing the cmd
		if err := execCmd.Run(); err != nil {
			return fmt.Errorf("failed to run cargo init for %q: %w", projectName, err)
		}

		if err := runFazInit(projectName); err != nil {
			return err
		}

		if err := appendGitInfoExclude(projectName); err != nil {
			return err
		}

		// copying justfile
		src := "assets/justfile_rust"
		out := filepath.Join(projectName, "justfile")

		if err := copyEmbFile(assets, src, out); err != nil {
			return err
		}

		fmt.Println("\nDone!!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rustCmd)
	rustCmd.Flags().BoolVar(&rsLibFlag, "lib", false, "Create a lib crate")
	// this does nothing. Adding just for completness
	rustCmd.Flags().Bool("bin", false, "Create a binary crate (default)")
}
