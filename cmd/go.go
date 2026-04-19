package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func newGoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "go [project_name]",
		Short: "Scaffold a go project",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]

			fmt.Println("Preparing your go project...")

			dirCmd := exec.Command(
				"mkdir",
				"-p",
				projectName,
			)

			if err := dirCmd.Run(); err != nil {
				return fmt.Errorf("failed to create project dir %w", err)
			}

			gitCmd := exec.Command(
				"git",
				"init",
			)
			gitCmd.Dir = projectName

			if err := gitCmd.Run(); err != nil {
				return fmt.Errorf("failed to run git init %w", err)
			}

			modCmd := exec.Command(
				"go",
				"mod",
				"init",
				projectName)
			modCmd.Dir = projectName

			if err := modCmd.Run(); err != nil {
				return fmt.Errorf("failed to run go mod init %w", err)
			}

			if err := runFazInit(projectName); err != nil {
				return err
			}

			if err := appendGitInfoExclude(projectName); err != nil {
				return err
			}

			src := "assets/justfile_go"
			out := filepath.Join(projectName, "justfile")

			if err := copyEmbFile(assets, src, out); err != nil {
				return err
			}

			src = "assets/main.go"
			out = filepath.Join(projectName, "main.go")

			if err := copyEmbFile(assets, src, out); err != nil {
				return err
			}

			fmt.Println("\nDone!!")

			return nil
		},
	}
}

// addGoCommand wires the Go scaffolding command into the root tree.
func addGoCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newGoCmd())
}
