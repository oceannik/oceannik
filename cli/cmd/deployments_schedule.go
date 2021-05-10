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

// scheduleCmd represents the schedule command
var deploymentsScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "schedule a new deployment",
	Run:   deploymentsScheduleCmdRun,
}

func deploymentsScheduleCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Scheduling new deployment...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetDeploymentServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.ScheduleDeploymentRequest{
		Namespace: namespace,
		Project:   args[0],
	}

	deployment, err := client.ScheduleDeployment(ctx, req)
	if err != nil {
		log.Fatalf("Could not schedule new deployment: %v", err)
	}

	table := utils.NewTable(os.Stdout, []string{"ID", "Namespace", "Project", "Status", "Scheduled at", "Started at", "Exited at"})

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
	table.Render()
}

func init() {
	deploymentsCmd.AddCommand(deploymentsScheduleCmd)
}
