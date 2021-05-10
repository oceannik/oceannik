package server

import (
	"context"
	"fmt"
	"log"

	"github.com/oceannik/oceannik/database"
	pb "github.com/oceannik/oceannik/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type DeploymentServiceServer struct {
	pb.UnimplementedDeploymentServiceServer
	db         *gorm.DB
	runnerChan chan uint
}

func (s *DeploymentServiceServer) ListDeployments(r *pb.ListDeploymentsRequest, stream pb.DeploymentService_ListDeploymentsServer) error {
	deployments, result := database.GetDeployments(s.db, r.GetNamespace())
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch deployments: %s", result.Error)
	}

	for _, deployment := range *deployments {
		res := &pb.Deployment{
			Identifier:  fmt.Sprint(deployment.ID),
			Namespace:   deployment.Namespace.Name,
			Project:     deployment.Project.Name,
			Status:      pb.Deployment_DeploymentStatus(pb.Deployment_DeploymentStatus_value[deployment.Status]),
			ScheduledAt: timestamppb.New(deployment.ScheduledAt),
			StartedAt:   timestamppb.New(deployment.StartedAt),
			ExitedAt:    timestamppb.New(deployment.ExitedAt),
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *DeploymentServiceServer) ScheduleDeployment(ctx context.Context, r *pb.ScheduleDeploymentRequest) (*pb.Deployment, error) {
	deployment, result := database.CreateDeployment(s.db, r.GetNamespace(), r.GetProject())
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not schedule a new deployment: %s", result.Error)
	}

	if deployment.Status == pb.Deployment_SCHEDULED.String() {
		log.Printf("Schedule deployment %d", deployment.ID)
		s.runnerChan <- deployment.ID
	}

	res := &pb.Deployment{
		Identifier:  fmt.Sprint(deployment.ID),
		Namespace:   deployment.Namespace.Name,
		Project:     deployment.Project.Name,
		Status:      pb.Deployment_DeploymentStatus(pb.Deployment_DeploymentStatus_value[deployment.Status]),
		ScheduledAt: timestamppb.New(deployment.ScheduledAt),
		StartedAt:   timestamppb.New(deployment.StartedAt),
		ExitedAt:    timestamppb.New(deployment.ExitedAt),
	}

	return res, nil
}
