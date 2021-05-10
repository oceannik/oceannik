package cmd

import (
	"github.com/spf13/cobra"
)

// namespacesCmd represents the namespaces command
var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "List, create and update Namespaces",
}

func init() {
	rootCmd.AddCommand(namespacesCmd)
}
