package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	"github.com/oceannik/oceannik/cli/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

var namespacesSetCmd = &cobra.Command{
	Use:   "set",
	Short: "create or update a namespace",
	Run:   namespacesSetCmdRun,
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
		Name:        args[0],
		Description: args[1],
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

	namespacesSetCmd.Flags().BoolP("force", "f", false, "force the entry to be updated or created, even if there's a conflict")
}
