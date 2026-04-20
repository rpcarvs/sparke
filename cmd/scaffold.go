package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type rustProjectType string

const (
	rustBinaryCrate  rustProjectType = "bin"
	rustLibraryCrate rustProjectType = "lib"
)

type pythonProjectType string

const (
	pythonAppProject     pythonProjectType = "app"
	pythonLibraryProject pythonProjectType = "lib"
	pythonPackageProject pythonProjectType = "package"
)

// scaffoldGoProject creates a Go project in the target directory.
func scaffoldGoProject(targetDir string, moduleName string) error {
	targetDir = resolveTargetDir(targetDir)
	moduleName = strings.TrimSpace(moduleName)
	if moduleName == "" {
		return fmt.Errorf("go module name cannot be empty")
	}

	fmt.Println("Preparing your go project...")

	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return fmt.Errorf("failed to create project dir %w", err)
	}

	gitCmd := exec.Command("git", "init")
	gitCmd.Dir = targetDir
	if err := gitCmd.Run(); err != nil {
		return fmt.Errorf("failed to run git init %w", err)
	}

	modCmd := exec.Command("go", "mod", "init", moduleName)
	modCmd.Dir = targetDir
	if err := modCmd.Run(); err != nil {
		return fmt.Errorf("failed to run go mod init %w", err)
	}

	if err := runFazInit(targetDir); err != nil {
		return err
	}

	if err := appendGitInfoExclude(targetDir); err != nil {
		return err
	}

	if err := copyEmbFile(assets, "assets/justfile_go", filepath.Join(targetDir, "justfile")); err != nil {
		return err
	}

	if err := copyEmbFile(assets, "assets/main.go", filepath.Join(targetDir, "main.go")); err != nil {
		return err
	}

	fmt.Println("\nDone!!")

	return nil
}

// scaffoldRustProject creates a Rust project in the target directory.
func scaffoldRustProject(targetDir string, projectType rustProjectType) error {
	targetDir = resolveTargetDir(targetDir)

	style := "--bin"
	if projectType == rustLibraryCrate {
		style = "--lib"
	}

	fmt.Println("Preparing your rust project...")

	execCmd := exec.Command("cargo", "init", targetDir, style)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("failed to run cargo init for %q: %w", targetDir, err)
	}

	if err := runFazInit(targetDir); err != nil {
		return err
	}

	if err := appendGitInfoExclude(targetDir); err != nil {
		return err
	}

	if err := copyEmbFile(assets, "assets/justfile_rust", filepath.Join(targetDir, "justfile")); err != nil {
		return err
	}

	fmt.Println("\nDone!!")

	return nil
}

// scaffoldPythonProject creates a Python project in the target directory.
func scaffoldPythonProject(targetDir string, projectType pythonProjectType) error {
	targetDir = resolveTargetDir(targetDir)

	style := "--app"
	switch projectType {
	case pythonLibraryProject:
		style = "--lib"
	case pythonPackageProject:
		style = "--package"
	}

	fmt.Println("Preparing your python project...")

	execCmd := exec.Command("uv", "init", targetDir, style)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	if err := execCmd.Run(); err != nil {
		return fmt.Errorf("failed to run uv init for %q: %w", targetDir, err)
	}

	if err := runFazInit(targetDir); err != nil {
		return err
	}

	if err := appendGitInfoExclude(targetDir); err != nil {
		return err
	}

	if err := copyEmbFile(assets, "assets/justfile_python", filepath.Join(targetDir, "justfile")); err != nil {
		return err
	}

	fmt.Println("\nDone!!")

	return nil
}

// resolveTargetDir maps an empty directory choice to the current directory.
func resolveTargetDir(targetDir string) string {
	trimmed := strings.TrimSpace(targetDir)
	if trimmed == "" {
		return "."
	}

	return trimmed
}
