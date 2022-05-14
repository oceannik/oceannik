package cmd

import (
	"github.com/spf13/cobra"
)

// secretsCmd represents the secrets command
var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "List, create and update Secrets in Namespaces",
}

func init() {
	rootCmd.AddCommand(secretsCmd)
}
