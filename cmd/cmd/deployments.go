package cmd

import (
	"github.com/spf13/cobra"
)

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "List and schedule new Deployments",
}

func init() {
	rootCmd.AddCommand(deploymentsCmd)
}
