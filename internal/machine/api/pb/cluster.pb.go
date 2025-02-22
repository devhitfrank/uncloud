// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: internal/machine/api/pb/cluster.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MachineMember_MembershipState int32

const (
	MachineMember_UNKNOWN MachineMember_MembershipState = 0
	// The member is active.
	MachineMember_UP MachineMember_MembershipState = 1
	// The member is active, but at least one cluster member suspects its down. For all purposes,
	// a SUSPECT member is treated as if it were UP until either it refutes the suspicion (becoming UP)
	// or fails to do so (being declared DOWN).
	MachineMember_SUSPECT MachineMember_MembershipState = 2
	// The member is confirmed DOWN.
	MachineMember_DOWN MachineMember_MembershipState = 3
)

// Enum value maps for MachineMember_MembershipState.
var (
	MachineMember_MembershipState_name = map[int32]string{
		0: "UNKNOWN",
		1: "UP",
		2: "SUSPECT",
		3: "DOWN",
	}
	MachineMember_MembershipState_value = map[string]int32{
		"UNKNOWN": 0,
		"UP":      1,
		"SUSPECT": 2,
		"DOWN":    3,
	}
)

func (x MachineMember_MembershipState) Enum() *MachineMember_MembershipState {
	p := new(MachineMember_MembershipState)
	*p = x
	return p
}

func (x MachineMember_MembershipState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MachineMember_MembershipState) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_machine_api_pb_cluster_proto_enumTypes[0].Descriptor()
}

func (MachineMember_MembershipState) Type() protoreflect.EnumType {
	return &file_internal_machine_api_pb_cluster_proto_enumTypes[0]
}

func (x MachineMember_MembershipState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MachineMember_MembershipState.Descriptor instead.
func (MachineMember_MembershipState) EnumDescriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{2, 0}
}

type DNSRecord_RecordType int32

const (
	DNSRecord_UNSPECIFIED DNSRecord_RecordType = 0
	DNSRecord_A           DNSRecord_RecordType = 1
	DNSRecord_AAAA        DNSRecord_RecordType = 2
)

// Enum value maps for DNSRecord_RecordType.
var (
	DNSRecord_RecordType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "A",
		2: "AAAA",
	}
	DNSRecord_RecordType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"A":           1,
		"AAAA":        2,
	}
)

func (x DNSRecord_RecordType) Enum() *DNSRecord_RecordType {
	p := new(DNSRecord_RecordType)
	*p = x
	return p
}

func (x DNSRecord_RecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DNSRecord_RecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_machine_api_pb_cluster_proto_enumTypes[1].Descriptor()
}

func (DNSRecord_RecordType) Type() protoreflect.EnumType {
	return &file_internal_machine_api_pb_cluster_proto_enumTypes[1]
}

func (x DNSRecord_RecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DNSRecord_RecordType.Descriptor instead.
func (DNSRecord_RecordType) EnumDescriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{6, 0}
}

type AddMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Network *NetworkConfig `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *AddMachineRequest) Reset() {
	*x = AddMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMachineRequest) ProtoMessage() {}

func (x *AddMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMachineRequest.ProtoReflect.Descriptor instead.
func (*AddMachineRequest) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *AddMachineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddMachineRequest) GetNetwork() *NetworkConfig {
	if x != nil {
		return x.Network
	}
	return nil
}

type AddMachineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *MachineInfo `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *AddMachineResponse) Reset() {
	*x = AddMachineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMachineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMachineResponse) ProtoMessage() {}

func (x *AddMachineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMachineResponse.ProtoReflect.Descriptor instead.
func (*AddMachineResponse) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *AddMachineResponse) GetMachine() *MachineInfo {
	if x != nil {
		return x.Machine
	}
	return nil
}

type MachineMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *MachineInfo                  `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
	State   MachineMember_MembershipState `protobuf:"varint,2,opt,name=state,proto3,enum=api.MachineMember_MembershipState" json:"state,omitempty"`
}

func (x *MachineMember) Reset() {
	*x = MachineMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineMember) ProtoMessage() {}

func (x *MachineMember) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineMember.ProtoReflect.Descriptor instead.
func (*MachineMember) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *MachineMember) GetMachine() *MachineInfo {
	if x != nil {
		return x.Machine
	}
	return nil
}

func (x *MachineMember) GetState() MachineMember_MembershipState {
	if x != nil {
		return x.State
	}
	return MachineMember_UNKNOWN
}

type ListMachinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*MachineMember `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *ListMachinesResponse) Reset() {
	*x = ListMachinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMachinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachinesResponse) ProtoMessage() {}

func (x *ListMachinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachinesResponse.ProtoReflect.Descriptor instead.
func (*ListMachinesResponse) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ListMachinesResponse) GetMachines() []*MachineMember {
	if x != nil {
		return x.Machines
	}
	return nil
}

type Domain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Domain) Reset() {
	*x = Domain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Domain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Domain) ProtoMessage() {}

func (x *Domain) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Domain.ProtoReflect.Descriptor instead.
func (*Domain) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *Domain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateDomainRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*DNSRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *UpdateDomainRecordsRequest) Reset() {
	*x = UpdateDomainRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDomainRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDomainRecordsRequest) ProtoMessage() {}

func (x *UpdateDomainRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDomainRecordsRequest.ProtoReflect.Descriptor instead.
func (*UpdateDomainRecordsRequest) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDomainRecordsRequest) GetRecords() []*DNSRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type DNSRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   DNSRecord_RecordType `protobuf:"varint,2,opt,name=type,proto3,enum=api.DNSRecord_RecordType" json:"type,omitempty"`
	Values []string             `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *DNSRecord) Reset() {
	*x = DNSRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSRecord) ProtoMessage() {}

func (x *DNSRecord) ProtoReflect() protoreflect.Message {
	mi := &file_internal_machine_api_pb_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSRecord.ProtoReflect.Descriptor instead.
func (*DNSRecord) Descriptor() ([]byte, []int) {
	return file_internal_machine_api_pb_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *DNSRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSRecord) GetType() DNSRecord_RecordType {
	if x != nil {
		return x.Type
	}
	return DNSRecord_UNSPECIFIED
}

func (x *DNSRecord) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_internal_machine_api_pb_cluster_proto protoreflect.FileDescriptor

var file_internal_machine_api_pb_cluster_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x40, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x53,
	0x50, 0x45, 0x43, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03,
	0x22, 0x46, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x08,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x1c, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x4e, 0x53, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x96,
	0x01, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x05, 0x0a, 0x01, 0x41, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x41, 0x41, 0x41, 0x41, 0x10, 0x02, 0x32, 0x8d, 0x02, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x64, 0x64, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x73, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x6b, 0x69,
	0x2f, 0x75, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_machine_api_pb_cluster_proto_rawDescOnce sync.Once
	file_internal_machine_api_pb_cluster_proto_rawDescData = file_internal_machine_api_pb_cluster_proto_rawDesc
)

func file_internal_machine_api_pb_cluster_proto_rawDescGZIP() []byte {
	file_internal_machine_api_pb_cluster_proto_rawDescOnce.Do(func() {
		file_internal_machine_api_pb_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_machine_api_pb_cluster_proto_rawDescData)
	})
	return file_internal_machine_api_pb_cluster_proto_rawDescData
}

var file_internal_machine_api_pb_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_machine_api_pb_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_machine_api_pb_cluster_proto_goTypes = []any{
	(MachineMember_MembershipState)(0), // 0: api.MachineMember.MembershipState
	(DNSRecord_RecordType)(0),          // 1: api.DNSRecord.RecordType
	(*AddMachineRequest)(nil),          // 2: api.AddMachineRequest
	(*AddMachineResponse)(nil),         // 3: api.AddMachineResponse
	(*MachineMember)(nil),              // 4: api.MachineMember
	(*ListMachinesResponse)(nil),       // 5: api.ListMachinesResponse
	(*Domain)(nil),                     // 6: api.Domain
	(*UpdateDomainRecordsRequest)(nil), // 7: api.UpdateDomainRecordsRequest
	(*DNSRecord)(nil),                  // 8: api.DNSRecord
	(*NetworkConfig)(nil),              // 9: api.NetworkConfig
	(*MachineInfo)(nil),                // 10: api.MachineInfo
	(*emptypb.Empty)(nil),              // 11: google.protobuf.Empty
}
var file_internal_machine_api_pb_cluster_proto_depIdxs = []int32{
	9,  // 0: api.AddMachineRequest.network:type_name -> api.NetworkConfig
	10, // 1: api.AddMachineResponse.machine:type_name -> api.MachineInfo
	10, // 2: api.MachineMember.machine:type_name -> api.MachineInfo
	0,  // 3: api.MachineMember.state:type_name -> api.MachineMember.MembershipState
	4,  // 4: api.ListMachinesResponse.machines:type_name -> api.MachineMember
	8,  // 5: api.UpdateDomainRecordsRequest.records:type_name -> api.DNSRecord
	1,  // 6: api.DNSRecord.type:type_name -> api.DNSRecord.RecordType
	2,  // 7: api.Cluster.AddMachine:input_type -> api.AddMachineRequest
	11, // 8: api.Cluster.ListMachines:input_type -> google.protobuf.Empty
	11, // 9: api.Cluster.GetDomain:input_type -> google.protobuf.Empty
	7,  // 10: api.Cluster.UpdateDomainRecords:input_type -> api.UpdateDomainRecordsRequest
	3,  // 11: api.Cluster.AddMachine:output_type -> api.AddMachineResponse
	5,  // 12: api.Cluster.ListMachines:output_type -> api.ListMachinesResponse
	6,  // 13: api.Cluster.GetDomain:output_type -> api.Domain
	11, // 14: api.Cluster.UpdateDomainRecords:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_internal_machine_api_pb_cluster_proto_init() }
func file_internal_machine_api_pb_cluster_proto_init() {
	if File_internal_machine_api_pb_cluster_proto != nil {
		return
	}
	file_internal_machine_api_pb_machine_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_machine_api_pb_cluster_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddMachineRequest); i {
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
		file_internal_machine_api_pb_cluster_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddMachineResponse); i {
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
		file_internal_machine_api_pb_cluster_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MachineMember); i {
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
		file_internal_machine_api_pb_cluster_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListMachinesResponse); i {
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
		file_internal_machine_api_pb_cluster_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Domain); i {
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
		file_internal_machine_api_pb_cluster_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateDomainRecordsRequest); i {
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
		file_internal_machine_api_pb_cluster_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DNSRecord); i {
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
			RawDescriptor: file_internal_machine_api_pb_cluster_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_machine_api_pb_cluster_proto_goTypes,
		DependencyIndexes: file_internal_machine_api_pb_cluster_proto_depIdxs,
		EnumInfos:         file_internal_machine_api_pb_cluster_proto_enumTypes,
		MessageInfos:      file_internal_machine_api_pb_cluster_proto_msgTypes,
	}.Build()
	File_internal_machine_api_pb_cluster_proto = out.File
	file_internal_machine_api_pb_cluster_proto_rawDesc = nil
	file_internal_machine_api_pb_cluster_proto_goTypes = nil
	file_internal_machine_api_pb_cluster_proto_depIdxs = nil
}
