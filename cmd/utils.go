package cmd

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func copyEmbFile(fs embed.FS, srcPath string, outPath string) error {
	data, err := fs.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("error reading embedded justfile")
	}

	if err := os.WriteFile(outPath, data, 0o644); err != nil {
		return fmt.Errorf("error writing justfile")
	}
	return nil
}

// runFazInit initializes faz in the generated project directory.
func runFazInit(projectName string) error {
	fazCmd := exec.Command(
		"faz",
		"init",
	)
	fazCmd.Dir = projectName

	if err := fazCmd.Run(); err != nil {
		return fmt.Errorf("failed to run faz init %w", err)
	}

	return nil
}

// appendGitInfoExclude appends local-only patterns to git exclude.
// It keeps helper files out of git tracking for generated projects.
func appendGitInfoExclude(projectName string) error {
	excludePath := filepath.Join(projectName, ".git", "info", "exclude")
	exclude, err := os.OpenFile(excludePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("failed to open git exclude file %w", err)
	}

	ignoreEntries := "\n.codex/\n.claude/\n*CLAUDE.md\n*AGENTS.md\n*PLAN.md\n"
	if _, err := exclude.WriteString(ignoreEntries); err != nil {
		_ = exclude.Close()
		return fmt.Errorf("failed to write git exclude file %w", err)
	}

	if err := exclude.Close(); err != nil {
		return fmt.Errorf("failed to close git exclude file %w", err)
	}

	return nil
}
