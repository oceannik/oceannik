package cmd

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	//go:embed version.txt
	versionFromFile string
	ReleaseVersion  string = strings.TrimSpace(versionFromFile)
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("[Ocean] Oceannik CLI application (ocean) version %s\n", ReleaseVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
