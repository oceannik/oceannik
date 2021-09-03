package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	"github.com/oceannik/oceannik/cli/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var secretsGetCmdEnableVerboseOutput bool
var secretsGetCmdSaveValueToFilePath string

var secretsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a list of secrets or a single secret",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this command accepts only one argument: the key of the secret to get")
		} else if len(args) == 1 {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
		return nil
	},
	Run: secretsGetCmdRun,
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
	table := utils.NewTable(os.Stdout, []string{"Key", "Value", "Description", "Kind"})

	if len(args) > 0 {
		// get a specific secret by key
		printedAny = secretsGetCmdGetSingle(client, ctx, table, args[0])
	} else {
		// get all secrets in the selected namespace
		printedAny = secretsGetCmdListAll(client, ctx, table)
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

func secretsGetCmdGetSingle(client pb.SecretServiceClient, ctx context.Context, table *tablewriter.Table, secretKey string) bool {
	req := &pb.GetSecretRequest{
		Namespace: namespace,
		Key:       secretKey,
	}

	secret, err := client.GetSecret(ctx, req)
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	secretValue := "<cut>"
	if secretsGetCmdEnableVerboseOutput {
		secretValue = secret.GetValue()
	}

	table.Append([]string{secret.GetKey(), secretValue, secret.GetDescription(), secret.GetKind().String()})

	if secretsGetCmdSaveValueToFilePath != "" {
		data := []byte(secret.GetValue())
		err := ioutil.WriteFile(secretsGetCmdSaveValueToFilePath, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	return true
}

func secretsGetCmdListAll(client pb.SecretServiceClient, ctx context.Context, table *tablewriter.Table) bool {
	req := &pb.ListSecretsRequest{
		Namespace: namespace,
	}

	stream, err := client.ListSecrets(ctx, req)
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	printedAny := false

	for {
		secret, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}

		printedAny = true
		secretValue := "<cut>"
		if secretsGetCmdEnableVerboseOutput {
			secretValue = secret.GetValue()
		}

		table.Append([]string{secret.GetKey(), secretValue, secret.GetDescription(), secret.GetKind().String()})
	}
	return printedAny
}
