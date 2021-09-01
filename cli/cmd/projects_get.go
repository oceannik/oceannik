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

var projectsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get a list of projects or a project by its name",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("this command accepts only one argument: the name of the project to get")
		} else if len(args) == 1 {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
		return nil
	},
	Run: projectsGetCmdRun,
}

func projectsGetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Getting projects...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetProjectServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := client.ListProjects(ctx, &pb.ListProjectsRequest{})
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	printedAny := false
	table := utils.NewTable(os.Stdout, []string{"Name", "Description", "Repository URL", "Config Path"})

	for {
		project, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}
		printedAny = true
		table.Append([]string{project.GetName(), project.GetDescription(), project.GetRepositoryUrl(), project.GetConfigPath()})
	}

	if !printedAny {
		log.Printf("[Ocean] No projects found!")
	} else {
		table.Render()
	}
}

func init() {
	projectsCmd.AddCommand(projectsGetCmd)
}
