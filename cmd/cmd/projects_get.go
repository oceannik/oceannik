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
	pb "github.com/oceannik/oceannik/proto"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

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

	printedAny := false
	table := utils.NewTable(os.Stdout, []string{"Name", "Description", "Repository URL", "Branch", "Config Path"})

	if len(args) > 0 {
		// get a specific project by name
		printedAny = projectsGetCmdGetSingle(client, ctx, table, args[0])
	} else {
		// get all projects in the selected namespace
		printedAny = projectsGetCmdListAll(client, ctx, table)
	}

	if !printedAny {
		log.Printf("[Ocean] No projects found! Create a new project with `ocean projects set`")
	} else {
		table.Render()
	}
}

func init() {
	projectsCmd.AddCommand(projectsGetCmd)
}

func projectsGetCmdGetSingle(client pb.ProjectServiceClient, ctx context.Context, table *tablewriter.Table, projectName string) bool {
	req := &pb.GetProjectRequest{
		Name: projectName,
	}

	project, err := client.GetProject(ctx, req)
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	table.Append(projectToTableRow(project))

	return true
}

func projectsGetCmdListAll(client pb.ProjectServiceClient, ctx context.Context, table *tablewriter.Table) bool {
	stream, err := client.ListProjects(ctx, &pb.ListProjectsRequest{})
	if err != nil {
		log.Fatalf("%v: %v", client, err)
	}

	printedAny := false

	for {
		project, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v: %v", client, err)
		}
		printedAny = true
		table.Append(projectToTableRow(project))
	}

	return printedAny
}

func projectToTableRow(project *pb.Project) []string {
	return []string{
		project.GetName(),
		project.GetDescription(),
		project.GetRepositoryUrl(),
		project.GetRepositoryBranch(),
		project.GetConfigPath(),
	}
}
