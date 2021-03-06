// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: proto/deployments.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Deployment_DeploymentStatus int32

const (
	Deployment_UNSPECIFIED    Deployment_DeploymentStatus = 0
	Deployment_SCHEDULED      Deployment_DeploymentStatus = 1
	Deployment_STARTED        Deployment_DeploymentStatus = 2
	Deployment_EXITED_SUCCESS Deployment_DeploymentStatus = 3
	Deployment_EXITED_FAILURE Deployment_DeploymentStatus = 4
)

// Enum value maps for Deployment_DeploymentStatus.
var (
	Deployment_DeploymentStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SCHEDULED",
		2: "STARTED",
		3: "EXITED_SUCCESS",
		4: "EXITED_FAILURE",
	}
	Deployment_DeploymentStatus_value = map[string]int32{
		"UNSPECIFIED":    0,
		"SCHEDULED":      1,
		"STARTED":        2,
		"EXITED_SUCCESS": 3,
		"EXITED_FAILURE": 4,
	}
)

func (x Deployment_DeploymentStatus) Enum() *Deployment_DeploymentStatus {
	p := new(Deployment_DeploymentStatus)
	*p = x
	return p
}

func (x Deployment_DeploymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Deployment_DeploymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_deployments_proto_enumTypes[0].Descriptor()
}

func (Deployment_DeploymentStatus) Type() protoreflect.EnumType {
	return &file_proto_deployments_proto_enumTypes[0]
}

func (x Deployment_DeploymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Deployment_DeploymentStatus.Descriptor instead.
func (Deployment_DeploymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{0, 0}
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier  string                      `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Namespace   string                      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Project     string                      `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Status      Deployment_DeploymentStatus `protobuf:"varint,4,opt,name=status,proto3,enum=oceannik.Deployment_DeploymentStatus" json:"status,omitempty"`
	ScheduledAt *timestamppb.Timestamp      `protobuf:"bytes,5,opt,name=scheduled_at,json=scheduledAt,proto3" json:"scheduled_at,omitempty"`
	StartedAt   *timestamppb.Timestamp      `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	ExitedAt    *timestamppb.Timestamp      `protobuf:"bytes,7,opt,name=exited_at,json=exitedAt,proto3" json:"exited_at,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deployments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deployments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *Deployment) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Deployment) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Deployment) GetStatus() Deployment_DeploymentStatus {
	if x != nil {
		return x.Status
	}
	return Deployment_UNSPECIFIED
}

func (x *Deployment) GetScheduledAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledAt
	}
	return nil
}

func (x *Deployment) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Deployment) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

type DeploymentLogChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk string `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *DeploymentLogChunk) Reset() {
	*x = DeploymentLogChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deployments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeploymentLogChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentLogChunk) ProtoMessage() {}

func (x *DeploymentLogChunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deployments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentLogChunk.ProtoReflect.Descriptor instead.
func (*DeploymentLogChunk) Descriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{1}
}

func (x *DeploymentLogChunk) GetChunk() string {
	if x != nil {
		return x.Chunk
	}
	return ""
}

type ListDeploymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListDeploymentsRequest) Reset() {
	*x = ListDeploymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deployments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeploymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeploymentsRequest) ProtoMessage() {}

func (x *ListDeploymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deployments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeploymentsRequest.ProtoReflect.Descriptor instead.
func (*ListDeploymentsRequest) Descriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{2}
}

func (x *ListDeploymentsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *GetDeploymentRequest) Reset() {
	*x = GetDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deployments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentRequest) ProtoMessage() {}

func (x *GetDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deployments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeploymentRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type ScheduleDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Project   string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *ScheduleDeploymentRequest) Reset() {
	*x = ScheduleDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deployments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleDeploymentRequest) ProtoMessage() {}

func (x *ScheduleDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deployments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleDeploymentRequest.ProtoReflect.Descriptor instead.
func (*ScheduleDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{4}
}

func (x *ScheduleDeploymentRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ScheduleDeploymentRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type GetDeploymentLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *GetDeploymentLogsRequest) Reset() {
	*x = GetDeploymentLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_deployments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentLogsRequest) ProtoMessage() {}

func (x *GetDeploymentLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_deployments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentLogsRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentLogsRequest) Descriptor() ([]byte, []int) {
	return file_proto_deployments_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeploymentLogsRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_proto_deployments_proto protoreflect.FileDescriptor

var file_proto_deployments_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6f, 0x63, 0x65, 0x61, 0x6e,
	0x6e, 0x69, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x03, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x67, 0x0a,
	0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x45, 0x58, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x58, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x55, 0x52, 0x45, 0x10, 0x04, 0x22, 0x2a, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x75,
	0x6e, 0x6b, 0x22, 0x36, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x36, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x22, 0x53, 0x0a, 0x19, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x32, 0xd9, 0x02, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x6f,
	0x63, 0x65, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x6f, 0x63, 0x65, 0x61,
	0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x63, 0x65, 0x61,
	0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x22, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x6e, 0x69,
	0x6b, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x63, 0x65,
	0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x12,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x6e,
	0x69, 0x6b, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2f, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x6e, 0x69, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_deployments_proto_rawDescOnce sync.Once
	file_proto_deployments_proto_rawDescData = file_proto_deployments_proto_rawDesc
)

func file_proto_deployments_proto_rawDescGZIP() []byte {
	file_proto_deployments_proto_rawDescOnce.Do(func() {
		file_proto_deployments_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_deployments_proto_rawDescData)
	})
	return file_proto_deployments_proto_rawDescData
}

var file_proto_deployments_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_deployments_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_deployments_proto_goTypes = []interface{}{
	(Deployment_DeploymentStatus)(0),  // 0: oceannik.Deployment.DeploymentStatus
	(*Deployment)(nil),                // 1: oceannik.Deployment
	(*DeploymentLogChunk)(nil),        // 2: oceannik.DeploymentLogChunk
	(*ListDeploymentsRequest)(nil),    // 3: oceannik.ListDeploymentsRequest
	(*GetDeploymentRequest)(nil),      // 4: oceannik.GetDeploymentRequest
	(*ScheduleDeploymentRequest)(nil), // 5: oceannik.ScheduleDeploymentRequest
	(*GetDeploymentLogsRequest)(nil),  // 6: oceannik.GetDeploymentLogsRequest
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_proto_deployments_proto_depIdxs = []int32{
	0, // 0: oceannik.Deployment.status:type_name -> oceannik.Deployment.DeploymentStatus
	7, // 1: oceannik.Deployment.scheduled_at:type_name -> google.protobuf.Timestamp
	7, // 2: oceannik.Deployment.started_at:type_name -> google.protobuf.Timestamp
	7, // 3: oceannik.Deployment.exited_at:type_name -> google.protobuf.Timestamp
	3, // 4: oceannik.DeploymentService.ListDeployments:input_type -> oceannik.ListDeploymentsRequest
	4, // 5: oceannik.DeploymentService.GetDeployment:input_type -> oceannik.GetDeploymentRequest
	6, // 6: oceannik.DeploymentService.GetDeploymentLogs:input_type -> oceannik.GetDeploymentLogsRequest
	5, // 7: oceannik.DeploymentService.ScheduleDeployment:input_type -> oceannik.ScheduleDeploymentRequest
	1, // 8: oceannik.DeploymentService.ListDeployments:output_type -> oceannik.Deployment
	1, // 9: oceannik.DeploymentService.GetDeployment:output_type -> oceannik.Deployment
	2, // 10: oceannik.DeploymentService.GetDeploymentLogs:output_type -> oceannik.DeploymentLogChunk
	1, // 11: oceannik.DeploymentService.ScheduleDeployment:output_type -> oceannik.Deployment
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_deployments_proto_init() }
func file_proto_deployments_proto_init() {
	if File_proto_deployments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_deployments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deployments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeploymentLogChunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deployments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeploymentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deployments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deployments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleDeploymentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_deployments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentLogsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_deployments_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_deployments_proto_goTypes,
		DependencyIndexes: file_proto_deployments_proto_depIdxs,
		EnumInfos:         file_proto_deployments_proto_enumTypes,
		MessageInfos:      file_proto_deployments_proto_msgTypes,
	}.Build()
	File_proto_deployments_proto = out.File
	file_proto_deployments_proto_rawDesc = nil
	file_proto_deployments_proto_goTypes = nil
	file_proto_deployments_proto_depIdxs = nil
}
