package cmd

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

var secretsPullCmdSaveFilesToDir string

var secretsPullCmd = &cobra.Command{
	Use:   "pull",
	Short: "pull secrets locally",
	Run:   secretsPullCmdRun,
}

func secretsPullCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Pulling secrets...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetSecretServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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

		if secret.GetKind() == pb.SecretKind_FILE {
			secretName := secret.GetKey()
			secretPath := fmt.Sprintf("%s/%s", secretsPullCmdSaveFilesToDir, secretName)
			data := []byte(secret.GetValue())
			err := ioutil.WriteFile(secretPath, data, 0644)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("[Ocean] Secret pulled to %s", secretPath)
		}
	}
}

func init() {
	secretsCmd.AddCommand(secretsPullCmd)

	secretsPullCmd.Flags().StringVar(&secretsPullCmdSaveFilesToDir, "output-dir", "pulled-secrets", "pull the secrets to the given directory")
}
