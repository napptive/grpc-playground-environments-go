// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: playground-environments/entities.proto

package grpc_playground_environments_go

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Environment defining a workspace for a user or group of users associated with an account.
type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EnvironmentId with the environment identifier.
	EnvironmentId string `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// AccountName with the account name.
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	// Name of the environment.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the purpose of this environment.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Environment) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Environment) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Environment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// EnvironmentListResponse with a list of selected environments.
type EnvironmentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Environments returned.
	Environments []*Environment `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *EnvironmentListResponse) Reset() {
	*x = EnvironmentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentListResponse) ProtoMessage() {}

func (x *EnvironmentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentListResponse.ProtoReflect.Descriptor instead.
func (*EnvironmentListResponse) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{1}
}

func (x *EnvironmentListResponse) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

// ListEnvironmentsRequest with information about the target account if the list is intended
// to be limited to a single account. The current selected account as specified in the JWT
// will be used otherwise.
type ListEnvironmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AccountName restrict if present the returned environments to the given account. Cannot be
	// used in conjunction with AccountId.
	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	// AccountId restricts if present the returned environmetns to the given account. Cannot be
	// used in conjunction with AccountName.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *ListEnvironmentsRequest) Reset() {
	*x = ListEnvironmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEnvironmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEnvironmentsRequest) ProtoMessage() {}

func (x *ListEnvironmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEnvironmentsRequest.ProtoReflect.Descriptor instead.
func (*ListEnvironmentsRequest) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{2}
}

func (x *ListEnvironmentsRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *ListEnvironmentsRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

// EnvironmentQuota message containing quota information about a given environment.
type EnvironmentQuota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CpuQuota with the maximum CPU quota available for the user in Kubernetes CPU units.
	CpuQuota float32 `protobuf:"fixed32,2,opt,name=cpu_quota,json=cpuQuota,proto3" json:"cpu_quota,omitempty"`
	// UsedCpu with the number of Kubernetes CPU units used by the running user applications.
	UsedCpu float32 `protobuf:"fixed32,3,opt,name=used_cpu,json=usedCpu,proto3" json:"used_cpu,omitempty"`
	// MemoryQuota with the maximum amount of memory available for the user in MB.
	MemoryQuota float32 `protobuf:"fixed32,4,opt,name=memory_quota,json=memoryQuota,proto3" json:"memory_quota,omitempty"`
	// UsedMemory with the memory used in MB by the running user applications.
	UsedMemory float32 `protobuf:"fixed32,5,opt,name=used_memory,json=usedMemory,proto3" json:"used_memory,omitempty"`
	// StorageQuota with the maximum amount of storage available for the user in in MB.
	StorageQuota float32 `protobuf:"fixed32,6,opt,name=storage_quota,json=storageQuota,proto3" json:"storage_quota,omitempty"`
	// UsedStorage with the storage used in MB by the running user applications.
	UsedStorage float32 `protobuf:"fixed32,7,opt,name=used_storage,json=usedStorage,proto3" json:"used_storage,omitempty"`
}

func (x *EnvironmentQuota) Reset() {
	*x = EnvironmentQuota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentQuota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentQuota) ProtoMessage() {}

func (x *EnvironmentQuota) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentQuota.ProtoReflect.Descriptor instead.
func (*EnvironmentQuota) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{3}
}

func (x *EnvironmentQuota) GetCpuQuota() float32 {
	if x != nil {
		return x.CpuQuota
	}
	return 0
}

func (x *EnvironmentQuota) GetUsedCpu() float32 {
	if x != nil {
		return x.UsedCpu
	}
	return 0
}

func (x *EnvironmentQuota) GetMemoryQuota() float32 {
	if x != nil {
		return x.MemoryQuota
	}
	return 0
}

func (x *EnvironmentQuota) GetUsedMemory() float32 {
	if x != nil {
		return x.UsedMemory
	}
	return 0
}

func (x *EnvironmentQuota) GetStorageQuota() float32 {
	if x != nil {
		return x.StorageQuota
	}
	return 0
}

func (x *EnvironmentQuota) GetUsedStorage() float32 {
	if x != nil {
		return x.UsedStorage
	}
	return 0
}

// EnvironmentQuotaResponse message containing quota information about a given environment.
type EnvironmentQuotaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EnvironmentId with the environment identifier.
	EnvironmentId string `protobuf:"bytes,1,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// AccountName with the account name.
	AccountName string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	// Name of the environment.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the purpose of this environment.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Quota used and available on this environment.
	Quota *EnvironmentQuota `protobuf:"bytes,5,opt,name=quota,proto3" json:"quota,omitempty"`
}

func (x *EnvironmentQuotaResponse) Reset() {
	*x = EnvironmentQuotaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentQuotaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentQuotaResponse) ProtoMessage() {}

func (x *EnvironmentQuotaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentQuotaResponse.ProtoReflect.Descriptor instead.
func (*EnvironmentQuotaResponse) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{4}
}

func (x *EnvironmentQuotaResponse) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *EnvironmentQuotaResponse) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *EnvironmentQuotaResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvironmentQuotaResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EnvironmentQuotaResponse) GetQuota() *EnvironmentQuota {
	if x != nil {
		return x.Quota
	}
	return nil
}

// EnvironmentSelector with information to determine the target environment. The current selected
// environment as specified in the JWT will be used otherwise.
// Keep in mind that we can select an environment through EnvironmentQualifiedName or
// through the combination of AccountId and EnvironmentId, being incompatible the use of the two selectors at
// the same time
type EnvironmentSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EnvironmentQualifiedName (EnvQN) contains both the account name and the environment name as
	// <account_name>/<env_name>.
	EnvironmentQualifiedName string `protobuf:"bytes,1,opt,name=environment_qualified_name,json=environmentQualifiedName,proto3" json:"environment_qualified_name,omitempty"`
	// AccountId with the account identifier.
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// EnvironmentId with the environment identifier
	EnvironmentId string `protobuf:"bytes,3,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
}

func (x *EnvironmentSelector) Reset() {
	*x = EnvironmentSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentSelector) ProtoMessage() {}

func (x *EnvironmentSelector) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentSelector.ProtoReflect.Descriptor instead.
func (*EnvironmentSelector) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{5}
}

func (x *EnvironmentSelector) GetEnvironmentQualifiedName() string {
	if x != nil {
		return x.EnvironmentQualifiedName
	}
	return ""
}

func (x *EnvironmentSelector) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *EnvironmentSelector) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

// EnvironmentUser with the information of a user with access to a given environment.
type EnvironmentUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Username with the user login.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Email.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Role of the user in the environment.
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *EnvironmentUser) Reset() {
	*x = EnvironmentUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentUser) ProtoMessage() {}

func (x *EnvironmentUser) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentUser.ProtoReflect.Descriptor instead.
func (*EnvironmentUser) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{6}
}

func (x *EnvironmentUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *EnvironmentUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EnvironmentUser) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// EnvironmentInfoResponse with complete information about an environment. Depending on the
// permission of the requesting user, some information may be omitted.
type EnvironmentInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Environment information.
	Environment *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	// Quota information.
	Quota *EnvironmentQuota `protobuf:"bytes,2,opt,name=quota,proto3" json:"quota,omitempty"`
	// Users with access to the environment.
	Users []*EnvironmentUser `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *EnvironmentInfoResponse) Reset() {
	*x = EnvironmentInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentInfoResponse) ProtoMessage() {}

func (x *EnvironmentInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentInfoResponse.ProtoReflect.Descriptor instead.
func (*EnvironmentInfoResponse) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{7}
}

func (x *EnvironmentInfoResponse) GetEnvironment() *Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

func (x *EnvironmentInfoResponse) GetQuota() *EnvironmentQuota {
	if x != nil {
		return x.Quota
	}
	return nil
}

func (x *EnvironmentInfoResponse) GetUsers() []*EnvironmentUser {
	if x != nil {
		return x.Users
	}
	return nil
}

// CreateEnvironmentRequest is the message to request the creation of a new environment
type CreateEnvironmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AccountId with the parent account identifier.
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Name with the pretty name of the environment for user facing operations.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the purpose of this environment.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// PrimaryEnvironment flag to know which is the primary/default environment for an account.
	SetAsDefaultEnvironment bool `protobuf:"varint,4,opt,name=set_as_default_environment,json=setAsDefaultEnvironment,proto3" json:"set_as_default_environment,omitempty"`
	// Environment Quota requested
	// CPU in Kubernetes CPU units.
	Cpu float32 `protobuf:"fixed32,5,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// RAM in MB.
	Ram float32 `protobuf:"fixed32,6,opt,name=ram,proto3" json:"ram,omitempty"`
	// EphemeralStorage in MB.
	Storage float32 `protobuf:"fixed32,7,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *CreateEnvironmentRequest) Reset() {
	*x = CreateEnvironmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playground_environments_entities_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEnvironmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEnvironmentRequest) ProtoMessage() {}

func (x *CreateEnvironmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playground_environments_entities_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEnvironmentRequest.ProtoReflect.Descriptor instead.
func (*CreateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return file_playground_environments_entities_proto_rawDescGZIP(), []int{8}
}

func (x *CreateEnvironmentRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEnvironmentRequest) GetSetAsDefaultEnvironment() bool {
	if x != nil {
		return x.SetAsDefaultEnvironment
	}
	return false
}

func (x *CreateEnvironmentRequest) GetCpu() float32 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *CreateEnvironmentRequest) GetRam() float32 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *CreateEnvironmentRequest) GetStorage() float32 {
	if x != nil {
		return x.Storage
	}
	return 0
}

var File_playground_environments_entities_proto protoreflect.FileDescriptor

var file_playground_environments_entities_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2d, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x63, 0x0a, 0x17, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x5b, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xd6, 0x01, 0x0a,
	0x10, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x70, 0x75, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x70, 0x75, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x19,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x64, 0x43, 0x70, 0x75, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x51, 0x75, 0x6f,
	0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x18, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x22, 0x99, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x1a, 0x65,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x18, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x57, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x17, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x05,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x3e, 0x0a,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0xfc, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x0a, 0x1a, 0x73, 0x65, 0x74, 0x5f, 0x61, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x73, 0x65, 0x74, 0x41, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x63, 0x70, 0x75,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x72,
	0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x55, 0x5a, 0x53,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x70, 0x70, 0x74,
	0x69, 0x76, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x2d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2d, 0x67, 0x6f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playground_environments_entities_proto_rawDescOnce sync.Once
	file_playground_environments_entities_proto_rawDescData = file_playground_environments_entities_proto_rawDesc
)

func file_playground_environments_entities_proto_rawDescGZIP() []byte {
	file_playground_environments_entities_proto_rawDescOnce.Do(func() {
		file_playground_environments_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_playground_environments_entities_proto_rawDescData)
	})
	return file_playground_environments_entities_proto_rawDescData
}

var file_playground_environments_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_playground_environments_entities_proto_goTypes = []interface{}{
	(*Environment)(nil),              // 0: playground_environments.Environment
	(*EnvironmentListResponse)(nil),  // 1: playground_environments.EnvironmentListResponse
	(*ListEnvironmentsRequest)(nil),  // 2: playground_environments.ListEnvironmentsRequest
	(*EnvironmentQuota)(nil),         // 3: playground_environments.EnvironmentQuota
	(*EnvironmentQuotaResponse)(nil), // 4: playground_environments.EnvironmentQuotaResponse
	(*EnvironmentSelector)(nil),      // 5: playground_environments.EnvironmentSelector
	(*EnvironmentUser)(nil),          // 6: playground_environments.EnvironmentUser
	(*EnvironmentInfoResponse)(nil),  // 7: playground_environments.EnvironmentInfoResponse
	(*CreateEnvironmentRequest)(nil), // 8: playground_environments.CreateEnvironmentRequest
}
var file_playground_environments_entities_proto_depIdxs = []int32{
	0, // 0: playground_environments.EnvironmentListResponse.environments:type_name -> playground_environments.Environment
	3, // 1: playground_environments.EnvironmentQuotaResponse.quota:type_name -> playground_environments.EnvironmentQuota
	0, // 2: playground_environments.EnvironmentInfoResponse.environment:type_name -> playground_environments.Environment
	3, // 3: playground_environments.EnvironmentInfoResponse.quota:type_name -> playground_environments.EnvironmentQuota
	6, // 4: playground_environments.EnvironmentInfoResponse.users:type_name -> playground_environments.EnvironmentUser
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_playground_environments_entities_proto_init() }
func file_playground_environments_entities_proto_init() {
	if File_playground_environments_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_playground_environments_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
		file_playground_environments_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentListResponse); i {
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
		file_playground_environments_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEnvironmentsRequest); i {
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
		file_playground_environments_entities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentQuota); i {
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
		file_playground_environments_entities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentQuotaResponse); i {
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
		file_playground_environments_entities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentSelector); i {
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
		file_playground_environments_entities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentUser); i {
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
		file_playground_environments_entities_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentInfoResponse); i {
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
		file_playground_environments_entities_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEnvironmentRequest); i {
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
			RawDescriptor: file_playground_environments_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_playground_environments_entities_proto_goTypes,
		DependencyIndexes: file_playground_environments_entities_proto_depIdxs,
		MessageInfos:      file_playground_environments_entities_proto_msgTypes,
	}.Build()
	File_playground_environments_entities_proto = out.File
	file_playground_environments_entities_proto_rawDesc = nil
	file_playground_environments_entities_proto_goTypes = nil
	file_playground_environments_entities_proto_depIdxs = nil
}
