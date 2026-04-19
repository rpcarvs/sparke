package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newPythonCmd() *cobra.Command {
	var (
		libFlag     bool
		packageFlag bool
	)

	pythonCmd := &cobra.Command{
		Use:   "python [project_name]",
		Short: "Scaffold a python project using uv",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			style := "--app"
			switch {
			case libFlag:
				style = "--lib"
			case packageFlag:
				style = "--package"
			}

			fmt.Println("Preparing your python project...")

			execCmd := exec.Command(
				"uv",
				"init",
				projectName,
				style,
			)

			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr

			if err := execCmd.Run(); err != nil {
				return fmt.Errorf("failed to run uv init for %q: %w", projectName, err)
			}

			if err := runFazInit(projectName); err != nil {
				return err
			}

			if err := appendGitInfoExclude(projectName); err != nil {
				return err
			}

			src := "assets/justfile_python"
			out := filepath.Join(projectName, "justfile")

			if err := copyEmbFile(assets, src, out); err != nil {
				return err
			}

			fmt.Println("\nDone!!")

			return nil
		},
	}

	pythonCmd.Flags().BoolVar(&libFlag, "lib", false, "Create a lib")
	pythonCmd.Flags().BoolVar(&packageFlag, "package", false, "Create a package")
	pythonCmd.Flags().Bool("app", false, "Create an app (default)")

	return pythonCmd
}

// addPythonCommand wires the Python scaffolding command into the root tree.
func addPythonCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newPythonCmd())
}
