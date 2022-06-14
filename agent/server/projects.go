package server

import (
	"context"
	"errors"

	"github.com/oceannik/oceannik/agent/database"
	pb "github.com/oceannik/oceannik/common/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProjectServiceServer struct {
	pb.UnimplementedProjectServiceServer
	db *gorm.DB
}

func projectAsProtobufStruct(project *database.Project) *pb.Project {
	return &pb.Project{
		Name:             project.Name,
		Description:      project.Description,
		RepositoryUrl:    project.RepositoryUrl,
		RepositoryBranch: project.RepositoryBranch,
		ConfigPath:       project.ConfigPath,
	}
}

func (s *ProjectServiceServer) ListProjects(r *pb.ListProjectsRequest, stream pb.ProjectService_ListProjectsServer) error {
	projects, result := database.GetProjects(s.db)
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch projects: %s", result.Error)
	}

	for _, project := range *projects {
		project := project
		res := projectAsProtobufStruct(&project)

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
}

func (s *ProjectServiceServer) GetProject(ctx context.Context, r *pb.GetProjectRequest) (*pb.Project, error) {
	project, result := database.GetProjectByName(s.db, r.GetName())
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not fetch project: %s", result.Error)
	}

	res := projectAsProtobufStruct(project)

	return res, nil
}

func (s *ProjectServiceServer) SetProject(ctx context.Context, r *pb.SetProjectRequest) (*pb.Project, error) {
	project, result := database.GetProjectByName(s.db, r.GetName())
	if result.Error != nil {
		// project does not exist, let's create a new one
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			project, result = database.CreateProject(s.db, r.GetName(), r.GetDescription(), r.GetRepositoryUrl(), r.GetRepositoryBranch(), r.GetConfigPath())
		}
	} else {
		// project does exist, do we have permission to overwrite it?
		if r.GetOverwriteIfExists() {
			project, result = database.UpdateProject(s.db, r.GetName(), r.GetDescription(), r.GetRepositoryUrl(), r.GetRepositoryBranch(), r.GetConfigPath())
		} else {
			return nil, status.Errorf(codes.Internal, "project with this name already exists")
		}
	}

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not create a new project: %s", result.Error)
	}

	res := projectAsProtobufStruct(project)

	return res, nil
}
