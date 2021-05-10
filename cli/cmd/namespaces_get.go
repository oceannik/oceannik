package cmd

import (
	"context"
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
	Short: "get a namespace or namespaces",
	Run:   namespacesGetCmdRun,
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
