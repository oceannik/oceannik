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

var projectsSetCmdDescription string
var projectsSetCmdForce bool

var projectsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "create or update a project",
	Run:   projectsSetCmdRun,
}

func projectsSetCmdRun(cmd *cobra.Command, args []string) {
	log.Printf("[Ocean] Setting a project...")

	agentConnector := connectors.AgentConnector{}
	agentConnector.Open()
	defer agentConnector.Close()

	client := agentConnector.GetProjectServiceClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SetProjectRequest{
		Name:              args[0],
		Description:       projectsSetCmdDescription,
		RepositoryUrl:     args[1],
		ConfigPath:        args[2],
		OverwriteIfExists: projectsSetCmdForce,
	}

	project, err := client.SetProject(ctx, req)
	if err != nil {
		log.Fatalf("Could not create new project: %v", err)
	}

	table := utils.NewTable(os.Stdout, []string{"Name", "Description", "Repository URL", "Config Path"})
	table.Append([]string{project.GetName(), project.GetDescription(), project.GetRepositoryUrl(), project.GetConfigPath()})
	table.Render()
}

func init() {
	projectsCmd.AddCommand(projectsSetCmd)

	projectsSetCmd.Flags().BoolVarP(&projectsSetCmdForce, "force", "f", false, "force the entry to be updated or created, even if there's a conflict")
	projectsSetCmd.Flags().StringVarP(&projectsSetCmdDescription, "description", "d", "", "set description for the project")

}
