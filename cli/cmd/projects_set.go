package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oceannik/oceannik/cli/connectors"
	"github.com/oceannik/oceannik/cli/utils"
	pb "github.com/oceannik/oceannik/proto"
	"github.com/spf13/cobra"
)

var projectsSetCmdDescription string
var projectsSetCmdRepositoryUrl string
var projectsSetCmdRepositoryBranch string
var projectsSetCmdConfigPath string
var projectsSetCmdForce bool

var projectsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "create or update a project",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("this command accepts only one argument: the name of the project to set")
		} else {
			if utils.IsValidIdentifierString(args[0]) {
				return nil
			}
			return fmt.Errorf("invalid characters used in argument: %s", args[0])
		}
	},
	Run: projectsSetCmdRun,
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
		OverwriteIfExists: projectsSetCmdForce,
	}

	if projectsSetCmdForce {
		// user wants to update the project
		if cmd.Flags().Lookup("description").Changed {
			req.Description = projectsSetCmdDescription
		}
		if cmd.Flags().Lookup("repository-url").Changed {
			req.RepositoryUrl = projectsSetCmdRepositoryUrl
		}
		if cmd.Flags().Lookup("repository-branch").Changed {
			req.RepositoryBranch = projectsSetCmdRepositoryBranch
		}
		if cmd.Flags().Lookup("oceannik-config").Changed {
			req.ConfigPath = projectsSetCmdConfigPath
		}
	} else {
		// user wants to create a new project
		req = &pb.SetProjectRequest{
			Name:              args[0],
			Description:       projectsSetCmdDescription,
			RepositoryUrl:     projectsSetCmdRepositoryUrl,
			RepositoryBranch:  projectsSetCmdRepositoryBranch,
			ConfigPath:        projectsSetCmdConfigPath,
			OverwriteIfExists: projectsSetCmdForce,
		}
	}

	project, err := client.SetProject(ctx, req)
	if err != nil {
		log.Fatalf("Could not create new project: %v", err)
	}

	table := utils.NewTable(os.Stdout, []string{"Name", "Description", "Repository URL", "Branch", "Config Path"})
	table.Append([]string{project.GetName(), project.GetDescription(), project.GetRepositoryUrl(), project.GetRepositoryBranch(), project.GetConfigPath()})
	table.Render()
}

func init() {
	projectsCmd.AddCommand(projectsSetCmd)

	projectsSetCmd.Flags().BoolVarP(&projectsSetCmdForce, "force", "f", false, "force the record to be updated")
	projectsSetCmd.Flags().StringVarP(&projectsSetCmdDescription, "description", "d", "", "set description for the project")
	projectsSetCmd.Flags().StringVarP(&projectsSetCmdRepositoryUrl, "repository-url", "r", "https://github.com/oceannik/examples", "set the repository url for the project")
	projectsSetCmd.Flags().StringVarP(&projectsSetCmdRepositoryBranch, "repository-branch", "b", "", "set the branch of the repository for the project")
	projectsSetCmd.Flags().StringVarP(&projectsSetCmdConfigPath, "oceannik-config", "o", "example-project/oceannik.yml", "set the path to the oceannik.yml service config file in the repository")
}
