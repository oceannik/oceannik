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

var deploymentsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a deployment or deployments",
	Run:   deploymentsGetCmdRun,
}

func deploymentsGetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Getting deployments...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetDeploymentServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.ListDeployments(ctx, &pb.ListDeploymentsRequest{Namespace: namespace})
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	printedAny := false
	table := utils.NewTable(os.Stdout, []string{"ID", "Namespace", "Project", "Status", "Scheduled at", "Started at", "Exited at"})

	for {
		deployment, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}
		printedAny = true

		startedAt := "<not started>"
		if !deployment.GetStartedAt().AsTime().IsZero() {
			startedAt = deployment.GetStartedAt().AsTime().Format(defaultTimeFormat)
		}
		exitedAt := "<not exited>"
		if !deployment.GetExitedAt().AsTime().IsZero() {
			exitedAt = deployment.GetExitedAt().AsTime().Format(defaultTimeFormat)
		}

		table.Append([]string{
			deployment.GetIdentifier(),
			deployment.GetNamespace(),
			deployment.GetProject(),
			deployment.GetStatus().String(),
			deployment.GetScheduledAt().AsTime().Format(defaultTimeFormat),
			startedAt,
			exitedAt,
		})
	}

	if !printedAny {
		log.Printf("[Ocean] No deployments found! Schedule a new deployment with `ocean deployments schedule`")
	} else {
		table.Render()
	}
}

func init() {
	deploymentsCmd.AddCommand(deploymentsGetCmd)
}
