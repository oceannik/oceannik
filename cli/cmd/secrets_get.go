package cmd

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	"github.com/oceannik/oceannik/cli/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

var secretsGetCmdEnableVerboseOutput bool
var secretsGetCmdSaveValueToFilePath string

// secretsGetCmd represents the get command
var secretsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a secret",
	Run:   secretsGetCmdRun,
}

// TODO: Implement output trimming for lines longer than 64
// strings.Split(s, char) + map + join

func secretsGetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Getting secrets...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetSecretServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	printedAny := false
	secretValue := "<cut>"
	table := utils.NewTable(os.Stdout, []string{"Namespace", "Key", "Value", "Description", "Kind"})

	if len(args) > 0 {
		// get a specific secret by key
		req := &pb.GetSecretRequest{
			Namespace: namespace,
			Key:       args[0],
		}

		secret, err := client.GetSecret(ctx, req)
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}

		printedAny = true

		if secretsGetCmdEnableVerboseOutput {
			secretValue = secret.GetValue()
		}

		table.Append([]string{secret.GetNamespace(), secret.GetKey(), secretValue, secret.GetDescription(), secret.GetKind().String()})

		if secretsGetCmdSaveValueToFilePath != "" {
			data := []byte(secret.GetValue())
			err := ioutil.WriteFile(secretsGetCmdSaveValueToFilePath, data, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		// get all secrets in the selected namespace
		req := &pb.ListSecretsRequest{
			Namespace: namespace,
		}

		stream, err := client.ListSecrets(ctx, req)
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}

		for {
			secret, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v: %v", client, err)
			}
			printedAny = true

			if secretsGetCmdEnableVerboseOutput {
				secretValue = secret.GetValue()
			}

			table.Append([]string{secret.GetNamespace(), secret.GetKey(), secretValue, secret.GetDescription(), secret.GetKind().String()})
		}
	}

	if !printedAny {
		log.Printf("[Ocean] No secrets found!")
	} else {
		table.Render()
	}
}

func init() {
	secretsCmd.AddCommand(secretsGetCmd)

	secretsGetCmd.Flags().BoolVarP(&secretsGetCmdEnableVerboseOutput, "verbose", "v", false, "display all fields (warning: this may expose your secrets!)")
	secretsGetCmd.Flags().StringVar(&secretsGetCmdSaveValueToFilePath, "value-to-file", "", "save the value of the secret to a file at the given location")

}
