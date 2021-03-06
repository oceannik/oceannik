// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (SecretService_ListSecretsClient, error)
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*Secret, error)
	SetSecret(ctx context.Context, in *SetSecretRequest, opts ...grpc.CallOption) (*Secret, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (SecretService_ListSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecretService_ServiceDesc.Streams[0], "/oceannik.SecretService/ListSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretServiceListSecretsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecretService_ListSecretsClient interface {
	Recv() (*Secret, error)
	grpc.ClientStream
}

type secretServiceListSecretsClient struct {
	grpc.ClientStream
}

func (x *secretServiceListSecretsClient) Recv() (*Secret, error) {
	m := new(Secret)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/oceannik.SecretService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) SetSecret(ctx context.Context, in *SetSecretRequest, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := c.cc.Invoke(ctx, "/oceannik.SecretService/SetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	ListSecrets(*ListSecretsRequest, SecretService_ListSecretsServer) error
	GetSecret(context.Context, *GetSecretRequest) (*Secret, error)
	SetSecret(context.Context, *SetSecretRequest) (*Secret, error)
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) ListSecrets(*ListSecretsRequest, SecretService_ListSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}
func (UnimplementedSecretServiceServer) GetSecret(context.Context, *GetSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretServiceServer) SetSecret(context.Context, *SetSecretRequest) (*Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSecret not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_ListSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSecretsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecretServiceServer).ListSecrets(m, &secretServiceListSecretsServer{stream})
}

type SecretService_ListSecretsServer interface {
	Send(*Secret) error
	grpc.ServerStream
}

type secretServiceListSecretsServer struct {
	grpc.ServerStream
}

func (x *secretServiceListSecretsServer) Send(m *Secret) error {
	return x.ServerStream.SendMsg(m)
}

func _SecretService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oceannik.SecretService/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_SetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).SetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oceannik.SecretService/SetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).SetSecret(ctx, req.(*SetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oceannik.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _SecretService_GetSecret_Handler,
		},
		{
			MethodName: "SetSecret",
			Handler:    _SecretService_SetSecret_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSecrets",
			Handler:       _SecretService_ListSecrets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/secrets.proto",
}
