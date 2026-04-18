package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	pyLibFlag     bool
	pyPackageFlag bool
)

var pythonCmd = &cobra.Command{
	Use:   "python [project_name]",
	Short: "Scaffold a python project using uv",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectName := args[0]
		// default
		style := "--app"
		switch {
		case pyLibFlag:
			style = "--lib"
		case pyPackageFlag:
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

		// executing the cmd
		if err := execCmd.Run(); err != nil {
			return fmt.Errorf("failed to run uv init for %q: %w", projectName, err)
		}

		if err := runFazInit(projectName); err != nil {
			return err
		}

		if err := appendGitInfoExclude(projectName); err != nil {
			return err
		}

		// copying justfile
		src := "assets/justfile_python"
		out := filepath.Join(projectName, "justfile")

		if err := copyEmbFile(assets, src, out); err != nil {
			return err
		}

		fmt.Println("\nDone!!")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)

	pythonCmd.Flags().BoolVar(&pyLibFlag, "lib", false, "Create a lib")
	pythonCmd.Flags().BoolVar(&pyPackageFlag, "package", false, "Create a package")
	// this does nothing. Adding just for completness
	pythonCmd.Flags().Bool("app", false, "Create an app (default)")
}
