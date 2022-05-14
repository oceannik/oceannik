package server

import (
	"context"
	"errors"

	"github.com/oceannik/oceannik/agent/database"
	pb "github.com/oceannik/oceannik/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type SecretServiceServer struct {
	pb.UnimplementedSecretServiceServer
	db *gorm.DB
}

func secretAsProtobufStruct(secret *database.Secret) *pb.Secret {
	return &pb.Secret{
		Namespace:   secret.Namespace.Name,
		Key:         secret.Key,
		Value:       secret.Value,
		Description: secret.Description,
		Kind:        pb.SecretKind(pb.SecretKind_value[secret.Kind]),
	}
}

func (s *SecretServiceServer) ListSecrets(r *pb.ListSecretsRequest, stream pb.SecretService_ListSecretsServer) error {
	secrets, result := database.GetSecrets(s.db, r.GetNamespace())
	if result.Error != nil {
		return status.Errorf(codes.Internal, "could not fetch secrets: %s", result.Error)
	}

	for _, secret := range *secrets {
		res := secretAsProtobufStruct(&secret)

		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func (s *SecretServiceServer) GetSecret(ctx context.Context, r *pb.GetSecretRequest) (*pb.Secret, error) {
	secret, result := database.GetSecretByKey(s.db, r.GetNamespace(), r.GetKey())
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not fetch secrets: %s", result.Error)
	}

	res := secretAsProtobufStruct(secret)

	return res, nil
}

func (s *SecretServiceServer) SetSecret(ctx context.Context, r *pb.SetSecretRequest) (*pb.Secret, error) {
	secret, result := database.GetSecretByKey(s.db, r.GetNamespace(), r.GetKey())
	if result.Error != nil {
		// secret does not exist, let's create a new one
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			secret, result = database.CreateSecret(s.db, r.GetNamespace(), r.GetKey(), r.GetValue(), r.GetDescription(), r.Kind.String())
		}
	} else {
		// secret does exist, do we have permission to overwrite it?
		if r.GetOverwriteIfExists() {
			secret, result = database.UpdateSecret(s.db, r.GetNamespace(), r.GetKey(), r.GetValue(), r.GetDescription(), r.Kind.String())
		} else {
			return nil, status.Errorf(codes.Internal, "secret with this key name already exists")
		}
	}

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "could not create a new secret: %s", result.Error)
	}

	res := secretAsProtobufStruct(secret)

	return res, nil
}
