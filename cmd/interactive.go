package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	huh "charm.land/huh/v2"
)

type projectLanguage string

const (
	goLanguage     projectLanguage = "go"
	rustLanguage   projectLanguage = "rust"
	pythonLanguage projectLanguage = "python"
)

// runInteractiveScaffold collects scaffold options from an interactive terminal.
func runInteractiveScaffold() error {
	var (
		language   projectLanguage
		rustType   rustProjectType
		pythonType pythonProjectType
		targetDir  string
		moduleName string
	)

	accessibleMode := os.Getenv("ACCESSIBLE") != ""

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[projectLanguage]().
				Title("Language").
				Description("Choose the project scaffold to generate.").
				Options(
					huh.NewOption("Go", goLanguage),
					huh.NewOption("Rust", rustLanguage),
					huh.NewOption("Python", pythonLanguage),
				).
				Value(&language),
		),
	).WithAccessible(accessibleMode).Run(); err != nil {
		return fmt.Errorf("failed to collect language selection %w", err)
	}

	switch language {
	case rustLanguage:
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[rustProjectType]().
					Title("Rust Project Type").
					Description("Choose the type of Rust crate to create.").
					Options(
						huh.NewOption("Binary Crate", rustBinaryCrate),
						huh.NewOption("Library Crate", rustLibraryCrate),
					).
					Value(&rustType),
			),
		).WithAccessible(accessibleMode).Run(); err != nil {
			return fmt.Errorf("failed to collect rust project type %w", err)
		}
	case pythonLanguage:
		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[pythonProjectType]().
					Title("Python Project Type").
					Description("Choose the type of Python project to create.").
					Options(
						huh.NewOption("App", pythonAppProject),
						huh.NewOption("Library", pythonLibraryProject),
						huh.NewOption("Package", pythonPackageProject),
					).
					Value(&pythonType),
			),
		).WithAccessible(accessibleMode).Run(); err != nil {
			return fmt.Errorf("failed to collect python project type %w", err)
		}
	}

	if err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Directory").
				Description("Leave blank to scaffold in the current directory.").
				Value(&targetDir),
		),
	).WithAccessible(accessibleMode).Run(); err != nil {
		return fmt.Errorf("failed to collect target directory %w", err)
	}

	if language == goLanguage {
		trimmedTargetDir := strings.TrimSpace(targetDir)
		if trimmedTargetDir == "" {
			moduleName = filepath.Base(currentWorkingDirectory())
		} else {
			moduleName = filepath.Base(trimmedTargetDir)
		}

		if err := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Go Module Name").
					Description("Set the module path for go mod init.").
					Validate(huh.ValidateNotEmpty()).
					Value(&moduleName),
			),
		).WithAccessible(accessibleMode).Run(); err != nil {
			return fmt.Errorf("failed to collect go module name %w", err)
		}
	}

	switch language {
	case goLanguage:
		return scaffoldGoProject(targetDir, moduleName)
	case rustLanguage:
		return scaffoldRustProject(targetDir, rustType)
	case pythonLanguage:
		return scaffoldPythonProject(targetDir, pythonType)
	default:
		return fmt.Errorf("unsupported project language %q", language)
	}
}

// isInteractiveTerminal reports whether standard input and output are terminals.
func isInteractiveTerminal() bool {
	return isTerminal(os.Stdin) && isTerminal(os.Stdout)
}

// isTerminal checks whether the provided file descriptor is a character device.
func isTerminal(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		return false
	}

	return (info.Mode() & os.ModeCharDevice) != 0
}

// currentWorkingDirectory returns the current directory or "." on lookup failure.
func currentWorkingDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		return "."
	}

	return cwd
}
