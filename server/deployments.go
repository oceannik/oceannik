package server

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

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

func deploymentAsProtobufStruct(deployment *database.Deployment) *pb.Deployment {
	return &pb.Deployment{
		Identifier:  fmt.Sprint(deployment.ID),
		Namespace:   deployment.Namespace.Name,
		Project:     deployment.Project.Name,
		Status:      pb.Deployment_DeploymentStatus(pb.Deployment_DeploymentStatus_value[deployment.Status]),
		ScheduledAt: timestamppb.New(deployment.ScheduledAt),
		StartedAt:   timestamppb.New(deployment.StartedAt),
		ExitedAt:    timestamppb.New(deployment.ExitedAt),
	}
}

func (s *DeploymentServiceServer) ListDeployments(r *pb.ListDeploymentsRequest, stream pb.DeploymentService_ListDeploymentsServer) error {
	deployments, result := database.GetDeployments(s.db, r.GetNamespace())
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch deployments: %s", result.Error)
	}

	for _, deployment := range *deployments {
		res := deploymentAsProtobufStruct(&deployment)

		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *DeploymentServiceServer) GetDeployment(ctx context.Context, r *pb.GetDeploymentRequest) (*pb.Deployment, error) {
	id, _ := strconv.Atoi(r.GetIdentifier())
	deployment, result := database.GetDeploymentByID(s.db, uint(id))
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not fetch deployment: %s", result.Error)
	}

	res := deploymentAsProtobufStruct(deployment)

	return res, nil
}

func (s *DeploymentServiceServer) ScheduleDeployment(ctx context.Context, r *pb.ScheduleDeploymentRequest) (*pb.Deployment, error) {
	deployment, result := database.CreateDeployment(s.db, r.GetNamespace(), r.GetProject())
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not schedule a new deployment: %s", result.Error)
	}

	if deployment.Status == pb.Deployment_SCHEDULED.String() {
		log.Printf("Schedule deployment %d", deployment.ID)
		select {
		case s.runnerChan <- deployment.ID:
		default:
			log.Printf("Channel is full, failing the deployment.")
			database.UpdateDeploymentStatus(s.db, deployment.ID, pb.Deployment_EXITED_FAILURE.String(), time.Time{}, time.Time{})
		}
	}

	res := deploymentAsProtobufStruct(deployment)

	return res, nil
}
