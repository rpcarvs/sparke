package cmd

import (
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
			projectType := pythonAppProject
			switch {
			case libFlag:
				projectType = pythonLibraryProject
			case packageFlag:
				projectType = pythonPackageProject
			}

			return scaffoldPythonProject(projectName, projectType)
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
