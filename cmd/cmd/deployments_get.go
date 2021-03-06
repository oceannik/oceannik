package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/client/connectors"
	"github.com/oceannik/oceannik/cmd/utils"
	pb "github.com/oceannik/oceannik/common/proto"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var deploymentsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a list of deployments or a specific deployment by its id",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this command accepts only one argument: the id of the deployment to get")
		} else if len(args) == 1 {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
		return nil
	},
	Run: deploymentsGetCmdRun,
}

func deploymentsGetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Getting deployments...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetDeploymentServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	printedAny := false
	table := utils.NewTable(os.Stdout, []string{"ID", "Namespace", "Project", "Status", "Scheduled at", "Started at", "Exited at"})

	if len(args) > 0 {
		// get a specific deployment by id
		printedAny = deploymentsGetCmdGetSingle(client, ctx, table, args[0])
	} else {
		// get all deployments in the selected namespace
		printedAny = deploymentsGetCmdListAll(client, ctx, table)
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

func deploymentsGetCmdGetSingle(client pb.DeploymentServiceClient, ctx context.Context, table *tablewriter.Table, deploymentId string) bool {
	req := &pb.GetDeploymentRequest{
		Identifier: deploymentId,
	}

	deployment, err := client.GetDeployment(ctx, req)
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	table.Append(deploymentToTableRow(deployment))

	return true
}

func deploymentsGetCmdListAll(client pb.DeploymentServiceClient, ctx context.Context, table *tablewriter.Table) bool {
	stream, err := client.ListDeployments(ctx, &pb.ListDeploymentsRequest{Namespace: namespace})
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	printedAny := false

	for {
		deployment, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}
		printedAny = true

		table.Append(deploymentToTableRow(deployment))
	}

	return printedAny
}

func deploymentToTableRow(deployment *pb.Deployment) []string {
	startedAt := "<not started>"
	if !deployment.GetStartedAt().AsTime().IsZero() {
		startedAt = deployment.GetStartedAt().AsTime().Format(defaultTimeFormat)
	}
	exitedAt := "<not exited>"
	if !deployment.GetExitedAt().AsTime().IsZero() {
		exitedAt = deployment.GetExitedAt().AsTime().Format(defaultTimeFormat)
	}

	return []string{
		deployment.GetIdentifier(),
		deployment.GetNamespace(),
		deployment.GetProject(),
		deployment.GetStatus().String(),
		deployment.GetScheduledAt().AsTime().Format(defaultTimeFormat),
		startedAt,
		exitedAt,
	}
}
