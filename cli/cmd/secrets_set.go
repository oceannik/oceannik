package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	"github.com/oceannik/oceannik/cli/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

var secretsSetCmdDescription string
var secretsSetCmdForce bool
var secretsSetCmdLoadValueFromFilePath string

// secretsSetCmd represents the set command
var secretsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "create or update a secret",
	Run:   secretsSetCmdRun,
}

func secretsSetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Setting a secret...")

	// log.Printf("Default timeout: %ds", defaultTimeoutInSeconds)
	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetSecretServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var secretValue string
	var secretKind pb.SecretKind

	if secretsSetCmdLoadValueFromFilePath != "" {
		b, err := ioutil.ReadFile(secretsSetCmdLoadValueFromFilePath)
		if err != nil {
			log.Fatal(err)
		}

		secretValue = string(b)
		secretKind = pb.SecretKind_FILE
	} else {
		secretValue = args[1]
		secretKind = pb.SecretKind_PLAIN
	}

	req := &pb.SetSecretRequest{
		Namespace:         namespace,
		Key:               args[0],
		Value:             secretValue,
		Description:       secretsSetCmdDescription,
		Kind:              secretKind,
		OverwriteIfExists: secretsSetCmdForce,
	}

	secret, err := client.SetSecret(ctx, req)
	if err != nil {
		log.Fatalf("Could not create new secret: %v", err)
	}

	table := utils.NewTable(os.Stdout, []string{"Namespace", "Key", "Value", "Description", "Kind"})
	table.Append([]string{secret.GetNamespace(), secret.GetKey(), secret.GetValue(), secret.GetDescription(), secret.GetKind().String()})
	table.Render()
}

func init() {
	secretsCmd.AddCommand(secretsSetCmd)

	secretsSetCmd.Flags().BoolVarP(&secretsSetCmdForce, "force", "f", false, "force the entry to be updated or created, even if there's a conflict")
	secretsSetCmd.Flags().StringVarP(&secretsSetCmdDescription, "description", "d", "", "set description for the secret")
	secretsSetCmd.Flags().StringVar(&secretsSetCmdLoadValueFromFilePath, "value-from-file", "", "load a file from a given path and use its contents as the value for the secret")
}
