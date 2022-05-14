package cmd

import (
	"log"

	server "github.com/oceannik/oceannik/agent/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmdEnableDevelopmentMode bool

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run an Agent server",
	Run: func(cmd *cobra.Command, args []string) {
		agentServerPort := viper.GetInt("agent.port")
		agentDebugServerHost := viper.GetString("agent.debug_server.host")
		agentDebugServerPort := viper.GetInt("agent.debug_server.port")
		agentDatabasePath := viper.GetString("agent.database_path")

		if serverCmdEnableDevelopmentMode {
			log.Print("[Oceannik Agent] Development mode is enabled.")
		}

		server.Start(agentServerPort, agentDatabasePath, serverCmdEnableDevelopmentMode, agentDebugServerHost, agentDebugServerPort)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().BoolVar(&serverCmdEnableDevelopmentMode, "dev", false, "run the server in development mode (insecure)")
}
