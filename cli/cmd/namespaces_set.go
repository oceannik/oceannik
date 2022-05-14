package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/utils"
	"github.com/oceannik/oceannik/client/connectors"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

var namespacesSetCmdDescription string
var namespacesSetCmdForce bool

var namespacesSetCmd = &cobra.Command{
	Use:   "set",
	Short: "create or update a namespace",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("this command accepts only one argument: the name of the namespace set")
		} else {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
	},
	Run: namespacesSetCmdRun,
}

func namespacesSetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Setting a namespace...")

	// log.Printf("Default timeout: %ds", defaultTimeoutInSeconds)
	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetNamespaceServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SetNamespaceRequest{
		Name:              args[0],
		OverwriteIfExists: namespacesSetCmdForce,
	}

	if namespacesSetCmdForce {
		// user wants to update the record
		if cmd.Flags().Lookup("description").Changed {
			req.Description = namespacesSetCmdDescription
		}
	} else {
		// user wants to create a new record
		req = &pb.SetNamespaceRequest{
			Name:              args[0],
			Description:       projectsSetCmdDescription,
			OverwriteIfExists: namespacesSetCmdForce,
		}
	}

	namespace, err := client.SetNamespace(ctx, req)
	if err != nil {
		log.Fatalf("Could not create new namespace: %v", err)
	}

	table := utils.NewTable(os.Stdout, []string{"Name", "Description", "Status"})
	table.Append([]string{namespace.GetName(), namespace.GetDescription(), "Created"})
	table.Render()
}

func init() {
	namespacesCmd.AddCommand(namespacesSetCmd)

	namespacesSetCmd.Flags().BoolVarP(&namespacesSetCmdForce, "force", "f", false, "force the record to be updated")
	namespacesSetCmd.Flags().StringVarP(&namespacesSetCmdDescription, "description", "d", "", "set description for the namespace")
}
