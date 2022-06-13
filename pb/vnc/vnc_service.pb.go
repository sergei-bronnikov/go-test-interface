// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: proto/vnc_service.proto

package vnc

import (
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

type GetStatusResponse_Status int32

const (
	GetStatusResponse_UNKOWN  GetStatusResponse_Status = 0
	GetStatusResponse_RUNNING GetStatusResponse_Status = 1
)

// Enum value maps for GetStatusResponse_Status.
var (
	GetStatusResponse_Status_name = map[int32]string{
		0: "UNKOWN",
		1: "RUNNING",
	}
	GetStatusResponse_Status_value = map[string]int32{
		"UNKOWN":  0,
		"RUNNING": 1,
	}
)

func (x GetStatusResponse_Status) Enum() *GetStatusResponse_Status {
	p := new(GetStatusResponse_Status)
	*p = x
	return p
}

func (x GetStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vnc_service_proto_enumTypes[0].Descriptor()
}

func (GetStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_proto_vnc_service_proto_enumTypes[0]
}

func (x GetStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetStatusResponse_Status.Descriptor instead.
func (GetStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{1, 0}
}

type GetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vnc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vnc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{0}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GetStatusResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=com.codio.chipmunk.proto.vnc.GetStatusResponse_Status" json:"status,omitempty"`
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vnc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vnc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatusResponse) GetStatus() GetStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetStatusResponse_UNKOWN
}

type CreateVncPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateVncPasswordRequest) Reset() {
	*x = CreateVncPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vnc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVncPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVncPasswordRequest) ProtoMessage() {}

func (x *CreateVncPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vnc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVncPasswordRequest.ProtoReflect.Descriptor instead.
func (*CreateVncPasswordRequest) Descriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{2}
}

type CreateVncPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateVncPasswordResponse) Reset() {
	*x = CreateVncPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vnc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVncPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVncPasswordResponse) ProtoMessage() {}

func (x *CreateVncPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vnc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVncPasswordResponse.ProtoReflect.Descriptor instead.
func (*CreateVncPasswordResponse) Descriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVncPasswordResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RestartVncServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartVncServiceRequest) Reset() {
	*x = RestartVncServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vnc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartVncServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartVncServiceRequest) ProtoMessage() {}

func (x *RestartVncServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vnc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartVncServiceRequest.ProtoReflect.Descriptor instead.
func (*RestartVncServiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{4}
}

type RestartVncServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartVncServiceResponse) Reset() {
	*x = RestartVncServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vnc_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartVncServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartVncServiceResponse) ProtoMessage() {}

func (x *RestartVncServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vnc_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartVncServiceResponse.ProtoReflect.Descriptor instead.
func (*RestartVncServiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_vnc_service_proto_rawDescGZIP(), []int{5}
}

var File_proto_vnc_service_proto protoreflect.FileDescriptor

var file_proto_vnc_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x64, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x6e, 0x63, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x69, 0x6f, 0x2e, 0x63, 0x68,
	0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x6e, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x21, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x4e, 0x4b, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6e,
	0x63, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x37, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6e, 0x63, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x56, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x56, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x81, 0x03, 0x0a, 0x03, 0x56, 0x6e, 0x63, 0x12, 0x6c, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x64, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x6e, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x64, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x6e, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x6e, 0x63, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x36,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x6e, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x6e, 0x63, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64,
	0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x6e, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6e, 0x63, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x84, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x6e, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x69,
	0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x6e, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x6e, 0x63, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x64, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75,
	0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x6e, 0x63, 0x2e, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x56, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x69, 0x6f, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x6e,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_vnc_service_proto_rawDescOnce sync.Once
	file_proto_vnc_service_proto_rawDescData = file_proto_vnc_service_proto_rawDesc
)

func file_proto_vnc_service_proto_rawDescGZIP() []byte {
	file_proto_vnc_service_proto_rawDescOnce.Do(func() {
		file_proto_vnc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vnc_service_proto_rawDescData)
	})
	return file_proto_vnc_service_proto_rawDescData
}

var file_proto_vnc_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_vnc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_vnc_service_proto_goTypes = []interface{}{
	(GetStatusResponse_Status)(0),     // 0: com.codio.chipmunk.proto.vnc.GetStatusResponse.Status
	(*GetStatusRequest)(nil),          // 1: com.codio.chipmunk.proto.vnc.GetStatusRequest
	(*GetStatusResponse)(nil),         // 2: com.codio.chipmunk.proto.vnc.GetStatusResponse
	(*CreateVncPasswordRequest)(nil),  // 3: com.codio.chipmunk.proto.vnc.CreateVncPasswordRequest
	(*CreateVncPasswordResponse)(nil), // 4: com.codio.chipmunk.proto.vnc.CreateVncPasswordResponse
	(*RestartVncServiceRequest)(nil),  // 5: com.codio.chipmunk.proto.vnc.RestartVncServiceRequest
	(*RestartVncServiceResponse)(nil), // 6: com.codio.chipmunk.proto.vnc.RestartVncServiceResponse
}
var file_proto_vnc_service_proto_depIdxs = []int32{
	0, // 0: com.codio.chipmunk.proto.vnc.GetStatusResponse.status:type_name -> com.codio.chipmunk.proto.vnc.GetStatusResponse.Status
	1, // 1: com.codio.chipmunk.proto.vnc.Vnc.GetStatus:input_type -> com.codio.chipmunk.proto.vnc.GetStatusRequest
	3, // 2: com.codio.chipmunk.proto.vnc.Vnc.CreateVncPassword:input_type -> com.codio.chipmunk.proto.vnc.CreateVncPasswordRequest
	5, // 3: com.codio.chipmunk.proto.vnc.Vnc.RestartVncService:input_type -> com.codio.chipmunk.proto.vnc.RestartVncServiceRequest
	2, // 4: com.codio.chipmunk.proto.vnc.Vnc.GetStatus:output_type -> com.codio.chipmunk.proto.vnc.GetStatusResponse
	4, // 5: com.codio.chipmunk.proto.vnc.Vnc.CreateVncPassword:output_type -> com.codio.chipmunk.proto.vnc.CreateVncPasswordResponse
	6, // 6: com.codio.chipmunk.proto.vnc.Vnc.RestartVncService:output_type -> com.codio.chipmunk.proto.vnc.RestartVncServiceResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_vnc_service_proto_init() }
func file_proto_vnc_service_proto_init() {
	if File_proto_vnc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vnc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusRequest); i {
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
		file_proto_vnc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResponse); i {
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
		file_proto_vnc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVncPasswordRequest); i {
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
		file_proto_vnc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVncPasswordResponse); i {
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
		file_proto_vnc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartVncServiceRequest); i {
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
		file_proto_vnc_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartVncServiceResponse); i {
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
			RawDescriptor: file_proto_vnc_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_vnc_service_proto_goTypes,
		DependencyIndexes: file_proto_vnc_service_proto_depIdxs,
		EnumInfos:         file_proto_vnc_service_proto_enumTypes,
		MessageInfos:      file_proto_vnc_service_proto_msgTypes,
	}.Build()
	File_proto_vnc_service_proto = out.File
	file_proto_vnc_service_proto_rawDesc = nil
	file_proto_vnc_service_proto_goTypes = nil
	file_proto_vnc_service_proto_depIdxs = nil
}
