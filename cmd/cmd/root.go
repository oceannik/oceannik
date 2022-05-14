package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/oceannik/oceannik/client/connectors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configDir  string
	namespace  string
	customHost string
	customPort int
	noTLS      bool
)

var (
	defaultTimeFormat = "Jan _2, 2006 15:04:05"
	agentConnector    = connectors.AgentConnector{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ocean",
	Short: "ocean: a CLI management tool for Oceannik instances.",
	Long: `ocean: a CLI management tool for Oceannik instances.

Visit https://github.com/oceannik/oceannik for more information.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().StringVarP(&configDir, "config-dir", "c", "", "config directory (default \"$HOME/.oceannik/\")")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "namespace to use for managing resources on the Agent")
	rootCmd.PersistentFlags().StringVar(&customHost, "host", "", "host to connect to/host to run the server on")
	rootCmd.PersistentFlags().IntVar(&customPort, "port", 5000, "port to connect to/port to run the server on")
	rootCmd.PersistentFlags().BoolVar(&noTLS, "no-tls", false, "disable TLS authentication")

	rootCmd.PersistentFlags()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	oceannikConfigRoot := fmt.Sprintf("%s/.oceannik", home)

	if configDir != "" {
		oceannikConfigRoot = configDir
	}

	oceannikConfigCerts := fmt.Sprintf("%s/certs", oceannikConfigRoot)

	viper.AddConfigPath(configDir)
	viper.SetConfigName(".ocean")
	viper.SetConfigType("yaml")

	viper.SetDefault("agent.host", "localhost")
	viper.SetDefault("agent.port", 5000)
	viper.SetDefault("agent.debug_server.host", "0.0.0.0")
	viper.SetDefault("agent.debug_server.port", 6060)
	viper.SetDefault("agent.debug_server.enable", false)
	viper.SetDefault("agent.deployments_queue_max_capacity", 42)
	viper.SetDefault("agent.database_path", fmt.Sprintf("%s/database.sqlite3", oceannikConfigRoot))
	viper.SetDefault("agent.runner.base_image", "ghcr.io/oceannik/runner-base-image:latest")
	viper.SetDefault("agent.runner.certs_path", fmt.Sprintf("%s/oceannik_runner", oceannikConfigCerts))

	viper.SetDefault("agent.certs.ca_cert_path", fmt.Sprintf("%s/oceannik_ca/oceannik_ca.crt", oceannikConfigCerts))
	viper.SetDefault("agent.certs.cert_path", fmt.Sprintf("%s/oceannik_agent.crt", oceannikConfigCerts))
	viper.SetDefault("agent.certs.key_path", fmt.Sprintf("%s/oceannik_agent.key", oceannikConfigCerts))

	viper.SetDefault("client.default_namespace", "default")
	viper.SetDefault("client.agent_host", "localhost")
	viper.SetDefault("client.agent_port", 5000)
	viper.SetDefault("client.dial_timeout", 5)

	viper.BindPFlag("agent.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("agent.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("client.default_namespace", rootCmd.PersistentFlags().Lookup("namespace"))
	viper.BindPFlag("client.agent_host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("client.agent_port", rootCmd.PersistentFlags().Lookup("port"))

	viper.SetDefault("client.certs.ca_cert_path", fmt.Sprintf("%s/oceannik_ca/oceannik_ca.crt", oceannikConfigCerts))
	viper.SetDefault("client.certs.cert_path", fmt.Sprintf("%s/oceannik_client.crt", oceannikConfigCerts))
	viper.SetDefault("client.certs.key_path", fmt.Sprintf("%s/oceannik_client.key", oceannikConfigCerts))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Println("[Ocean] Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println("[Ocean] No configuration file detected. Using provided defaults.")
	}
}
