package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/utils"
	"github.com/oceannik/oceannik/client/connectors"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 2 {
			return errors.New("this command accepts only two arguments: the key and the value of the secret")
		} else if len(args) == 2 {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		} else if len(args) == 1 {
			if !utils.IsValidIdentifierString(args[0]) {
				return fmt.Errorf("invalid characters used in argument: %s", args[0])
			}
			if !cmd.Flags().Lookup("value-from-file").Changed {
				return errors.New("if no value given as an argument, --value-from-file must be specified")
			}
			return nil
		}
		return errors.New("the key of the secret needs to be given as an argument")
	},
	Run: secretsSetCmdRun,
}

func secretsSetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Setting a secret...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetSecretServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var secretValue string
	var secretKind pb.SecretKind

	if secretsSetCmdLoadValueFromFilePath != "" {
		// user sets the value from file
		b, err := ioutil.ReadFile(secretsSetCmdLoadValueFromFilePath)
		if err != nil {
			log.Fatal(err)
		}

		secretValue = string(b)
		secretKind = pb.SecretKind_FILE
	} else if len(args) > 0 {
		// user sets the value from the command line
		secretValue = args[1]
		secretKind = pb.SecretKind_PLAIN
	}

	req := &pb.SetSecretRequest{
		Namespace:         namespace,
		Key:               args[0],
		OverwriteIfExists: secretsSetCmdForce,
	}

	if secretsSetCmdForce {
		// user wants to update the secret
		if secretValue != "" {
			req.Value = secretValue
			req.Kind = secretKind
		}
		if cmd.Flags().Lookup("description").Changed {
			req.Description = secretsSetCmdDescription
		}
	} else {
		// user wants to create a new secret
		req = &pb.SetSecretRequest{
			Namespace:         namespace,
			Key:               args[0],
			Value:             secretValue,
			Description:       secretsSetCmdDescription,
			Kind:              secretKind,
			OverwriteIfExists: secretsSetCmdForce,
		}
	}

	secret, err := client.SetSecret(ctx, req)
	if err != nil {
		log.Fatalf("Could not create new secret: %v", err)
	}

	table := utils.NewTable(os.Stdout, []string{"Key", "Value", "Description", "Kind"})
	table.Append([]string{secret.GetKey(), secret.GetValue(), secret.GetDescription(), secret.GetKind().String()})
	table.Render()
}

func init() {
	secretsCmd.AddCommand(secretsSetCmd)

	secretsSetCmd.Flags().BoolVarP(&secretsSetCmdForce, "force", "f", false, "force the record to be updated")
	secretsSetCmd.Flags().StringVarP(&secretsSetCmdDescription, "description", "d", "", "set description for the secret")
	secretsSetCmd.Flags().StringVar(&secretsSetCmdLoadValueFromFilePath, "value-from-file", "", "load a file from a given path and use its contents as the value for the secret")
}
