package server

import (
	"context"

	"github.com/oceannik/oceannik/database"
	pb "github.com/oceannik/oceannik/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type NamespaceServiceServer struct {
	pb.UnimplementedNamespaceServiceServer
	db *gorm.DB
}

func (s *NamespaceServiceServer) ListNamespaces(r *pb.ListNamespacesRequest, stream pb.NamespaceService_ListNamespacesServer) error {
	namespaces, result := database.GetNamespaces(s.db)
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch namespaces: %s", result.Error)
	}

	for _, namespace := range *namespaces {
		res := &pb.Namespace{
			Name:        namespace.Name,
			Description: namespace.Description,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *NamespaceServiceServer) GetNamespace(ctx context.Context, r *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}

func (s *NamespaceServiceServer) SetNamespace(ctx context.Context, r *pb.SetNamespaceRequest) (*pb.Namespace, error) {
	namespace, result := database.CreateNamespace(s.db, r.Name, r.Description)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not create new namespace: %s", result.Error)
	}

	res := &pb.Namespace{
		Name:        namespace.Name,
		Description: namespace.Description,
	}

	return res, nil
}
