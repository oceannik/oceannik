package cmd

import (
	"github.com/spf13/cobra"
)

// namespacesCmd represents the namespaces command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List, create and update Projects",
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
