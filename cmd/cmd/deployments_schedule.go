package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/client/connectors"
	"github.com/oceannik/oceannik/cmd/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

// scheduleCmd represents the schedule command
var deploymentsScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "schedule a new deployment",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this command accepts only one argument: the name of the project to deploy")
		} else if len(args) == 1 {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
		return errors.New("the name of the project to deploy needs to be given")
	},
	Run: deploymentsScheduleCmdRun,
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
