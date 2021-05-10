package cmd

import (
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var deploymentsLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "get logs of a deployment",
	Run:   deploymentsLogsCmdRun,
}

func deploymentsLogsCmdRun(cmd *cobra.Command, args []string) {}

func init() {
	deploymentsCmd.AddCommand(deploymentsLogsCmd)
}
