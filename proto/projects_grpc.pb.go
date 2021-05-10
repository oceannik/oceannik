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

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (ProjectService_ListProjectsClient, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	SetProject(ctx context.Context, in *SetProjectRequest, opts ...grpc.CallOption) (*Project, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (ProjectService_ListProjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[0], "/oceannik.ProjectService/ListProjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceListProjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_ListProjectsClient interface {
	Recv() (*Project, error)
	grpc.ClientStream
}

type projectServiceListProjectsClient struct {
	grpc.ClientStream
}

func (x *projectServiceListProjectsClient) Recv() (*Project, error) {
	m := new(Project)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/oceannik.ProjectService/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) SetProject(ctx context.Context, in *SetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/oceannik.ProjectService/SetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	ListProjects(*ListProjectsRequest, ProjectService_ListProjectsServer) error
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	SetProject(context.Context, *SetProjectRequest) (*Project, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) ListProjects(*ListProjectsRequest, ProjectService_ListProjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedProjectServiceServer) GetProject(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServiceServer) SetProject(context.Context, *SetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProject not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_ListProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListProjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).ListProjects(m, &projectServiceListProjectsServer{stream})
}

type ProjectService_ListProjectsServer interface {
	Send(*Project) error
	grpc.ServerStream
}

type projectServiceListProjectsServer struct {
	grpc.ServerStream
}

func (x *projectServiceListProjectsServer) Send(m *Project) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oceannik.ProjectService/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_SetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).SetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oceannik.ProjectService/SetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).SetProject(ctx, req.(*SetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oceannik.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "SetProject",
			Handler:    _ProjectService_SetProject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProjects",
			Handler:       _ProjectService_ListProjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/projects.proto",
}
