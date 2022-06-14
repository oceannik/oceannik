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

type NamespaceServiceServer struct {
	pb.UnimplementedNamespaceServiceServer
	db *gorm.DB
}

func namespaceAsProtobufStruct(namespace *database.Namespace) *pb.Namespace {
	return &pb.Namespace{
		Name:        namespace.Name,
		Description: namespace.Description,
	}
}

func (s *NamespaceServiceServer) ListNamespaces(r *pb.ListNamespacesRequest, stream pb.NamespaceService_ListNamespacesServer) error {
	namespaces, result := database.GetNamespaces(s.db)
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch namespaces: %s", result.Error)
	}

	for _, namespace := range *namespaces {
		namespace := namespace
		res := namespaceAsProtobufStruct(&namespace)

		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *NamespaceServiceServer) GetNamespace(ctx context.Context, r *pb.GetNamespaceRequest) (*pb.Namespace, error) {
	namespace, result := database.GetNamespaceByName(s.db, r.GetName())
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not fetch namespace: %s", result.Error)
	}

	res := namespaceAsProtobufStruct(namespace)

	return res, nil
}

func (s *NamespaceServiceServer) SetNamespace(ctx context.Context, r *pb.SetNamespaceRequest) (*pb.Namespace, error) {
	namespace, result := database.GetNamespaceByName(s.db, r.GetName())
	if result.Error != nil {
		// namespace does not exist, let's create a new one
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			namespace, result = database.CreateNamespace(s.db, r.GetName(), r.GetDescription())
		}
	} else {
		// namespace does exist, do we have permission to overwrite it?
		if r.GetOverwriteIfExists() {
			namespace, result = database.UpdateNamespace(s.db, r.GetName(), r.GetDescription())
		} else {
			return nil, status.Errorf(codes.Internal, "namespace with this name already exists")
		}
	}

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not create new namespace: %s", result.Error)
	}

	res := namespaceAsProtobufStruct(namespace)

	return res, nil
}
