package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	"github.com/oceannik/oceannik/cli/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

// var defaultTimeoutInSeconds = 5

var namespacesGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a list of namespaces or a namespace by its name",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this command accepts only one argument: the name of the namespace to get")
		} else if len(args) == 1 {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
		return nil
	},
	Run: namespacesGetCmdRun,
}

func namespacesGetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Getting namespaces...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetNamespaceServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.ListNamespaces(ctx, &pb.ListNamespacesRequest{})
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	printedAny := false
	table := utils.NewTable(os.Stdout, []string{"Name", "Description"})

	for {
		namespace, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}
		printedAny = true
		table.Append([]string{namespace.GetName(), namespace.GetDescription()})
	}

	if !printedAny {
		log.Printf("[Ocean] No namespaces found! Create a new namespace with `ocean namespaces set`")
	} else {
		table.Render()
	}
}

func init() {
	namespacesCmd.AddCommand(namespacesGetCmd)
}
