// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: permission_messages.proto

package authservicev1

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

type PermissionCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission *Permission `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *PermissionCreateRequest) Reset() {
	*x = PermissionCreateRequest{}
	mi := &file_permission_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionCreateRequest) ProtoMessage() {}

func (x *PermissionCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionCreateRequest.ProtoReflect.Descriptor instead.
func (*PermissionCreateRequest) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionCreateRequest) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type PermissionCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionCreateResponse) Reset() {
	*x = PermissionCreateResponse{}
	mi := &file_permission_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionCreateResponse) ProtoMessage() {}

func (x *PermissionCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionCreateResponse.ProtoReflect.Descriptor instead.
func (*PermissionCreateResponse) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PermissionReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionReadRequest) Reset() {
	*x = PermissionReadRequest{}
	mi := &file_permission_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionReadRequest) ProtoMessage() {}

func (x *PermissionReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionReadRequest.ProtoReflect.Descriptor instead.
func (*PermissionReadRequest) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PermissionReadRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PermissionReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission *Permission `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *PermissionReadResponse) Reset() {
	*x = PermissionReadResponse{}
	mi := &file_permission_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionReadResponse) ProtoMessage() {}

func (x *PermissionReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionReadResponse.ProtoReflect.Descriptor instead.
func (*PermissionReadResponse) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PermissionReadResponse) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type PermissionUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Permission *Permission `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *PermissionUpdateRequest) Reset() {
	*x = PermissionUpdateRequest{}
	mi := &file_permission_messages_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUpdateRequest) ProtoMessage() {}

func (x *PermissionUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUpdateRequest.ProtoReflect.Descriptor instead.
func (*PermissionUpdateRequest) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{4}
}

func (x *PermissionUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PermissionUpdateRequest) GetPermission() *Permission {
	if x != nil {
		return x.Permission
	}
	return nil
}

type PermissionUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionUpdateResponse) Reset() {
	*x = PermissionUpdateResponse{}
	mi := &file_permission_messages_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionUpdateResponse) ProtoMessage() {}

func (x *PermissionUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionUpdateResponse.ProtoReflect.Descriptor instead.
func (*PermissionUpdateResponse) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{5}
}

func (x *PermissionUpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PermissionDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PermissionDeleteRequest) Reset() {
	*x = PermissionDeleteRequest{}
	mi := &file_permission_messages_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDeleteRequest) ProtoMessage() {}

func (x *PermissionDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDeleteRequest.ProtoReflect.Descriptor instead.
func (*PermissionDeleteRequest) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{6}
}

func (x *PermissionDeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PermissionDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PermissionDeleteResponse) Reset() {
	*x = PermissionDeleteResponse{}
	mi := &file_permission_messages_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDeleteResponse) ProtoMessage() {}

func (x *PermissionDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDeleteResponse.ProtoReflect.Descriptor instead.
func (*PermissionDeleteResponse) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{7}
}

var File_permission_messages_proto protoreflect.FileDescriptor

var file_permission_messages_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x17, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x2a, 0x0a, 0x18, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x16, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x18, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22,
	0x5a, 0x20, 0x61, 0x70, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_messages_proto_rawDescOnce sync.Once
	file_permission_messages_proto_rawDescData = file_permission_messages_proto_rawDesc
)

func file_permission_messages_proto_rawDescGZIP() []byte {
	file_permission_messages_proto_rawDescOnce.Do(func() {
		file_permission_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_messages_proto_rawDescData)
	})
	return file_permission_messages_proto_rawDescData
}

var file_permission_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_permission_messages_proto_goTypes = []any{
	(*PermissionCreateRequest)(nil),  // 0: authservice.PermissionCreateRequest
	(*PermissionCreateResponse)(nil), // 1: authservice.PermissionCreateResponse
	(*PermissionReadRequest)(nil),    // 2: authservice.PermissionReadRequest
	(*PermissionReadResponse)(nil),   // 3: authservice.PermissionReadResponse
	(*PermissionUpdateRequest)(nil),  // 4: authservice.PermissionUpdateRequest
	(*PermissionUpdateResponse)(nil), // 5: authservice.PermissionUpdateResponse
	(*PermissionDeleteRequest)(nil),  // 6: authservice.PermissionDeleteRequest
	(*PermissionDeleteResponse)(nil), // 7: authservice.PermissionDeleteResponse
	(*Permission)(nil),               // 8: authservice.Permission
}
var file_permission_messages_proto_depIdxs = []int32{
	8, // 0: authservice.PermissionCreateRequest.permission:type_name -> authservice.Permission
	8, // 1: authservice.PermissionReadResponse.permission:type_name -> authservice.Permission
	8, // 2: authservice.PermissionUpdateRequest.permission:type_name -> authservice.Permission
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_permission_messages_proto_init() }
func file_permission_messages_proto_init() {
	if File_permission_messages_proto != nil {
		return
	}
	file_auth_models_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_permission_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_permission_messages_proto_goTypes,
		DependencyIndexes: file_permission_messages_proto_depIdxs,
		MessageInfos:      file_permission_messages_proto_msgTypes,
	}.Build()
	File_permission_messages_proto = out.File
	file_permission_messages_proto_rawDesc = nil
	file_permission_messages_proto_goTypes = nil
	file_permission_messages_proto_depIdxs = nil
}