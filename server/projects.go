package server

import (
	"context"

	"github.com/oceannik/oceannik/database"
	pb "github.com/oceannik/oceannik/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProjectServiceServer struct {
	pb.UnimplementedProjectServiceServer
	db *gorm.DB
}

func (s *ProjectServiceServer) ListProjects(r *pb.ListProjectsRequest, stream pb.ProjectService_ListProjectsServer) error {
	projects, result := database.GetProjects(s.db)
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch projects: %s", result.Error)
	}

	for _, project := range *projects {
		res := &pb.Project{
			Name:          project.Name,
			Description:   project.Description,
			RepositoryUrl: project.RepositoryUrl,
			ConfigPath:    project.ConfigPath,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *ProjectServiceServer) GetProject(ctx context.Context, r *pb.GetProjectRequest) (*pb.Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}

func (s *ProjectServiceServer) SetProject(ctx context.Context, r *pb.SetProjectRequest) (*pb.Project, error) {
	project, result := database.CreateProject(s.db, r.Name, r.Description, r.RepositoryUrl, r.ConfigPath)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not create new project: %s", result.Error)
	}

	res := &pb.Project{
		Name:          project.Name,
		Description:   project.Description,
		RepositoryUrl: project.RepositoryUrl,
		ConfigPath:    project.ConfigPath,
	}

	return res, nil
}
