// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/pb/userproto/user.proto

package userproto

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

type CreatUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	HouseName   string `protobuf:"bytes,2,opt,name=HouseName,proto3" json:"HouseName,omitempty"`
	City        string `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNumber int64  `protobuf:"varint,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
}

func (x *CreatUserRequest) Reset() {
	*x = CreatUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatUserRequest) ProtoMessage() {}

func (x *CreatUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatUserRequest.ProtoReflect.Descriptor instead.
func (*CreatUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreatUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatUserRequest) GetHouseName() string {
	if x != nil {
		return x.HouseName
	}
	return ""
}

func (x *CreatUserRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreatUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreatUserRequest) GetPhoneNumber() int64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	HouseName   string `protobuf:"bytes,3,opt,name=HouseName,proto3" json:"HouseName,omitempty"`
	City        string `protobuf:"bytes,4,opt,name=City,proto3" json:"City,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNumber int64  `protobuf:"varint,6,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetHouseName() string {
	if x != nil {
		return x.HouseName
	}
	return ""
}

func (x *UpdateUserRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPhoneNumber() int64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteUserRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{4}
}

func (x *CommonResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CommonResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetUserResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	HouseName   string `protobuf:"bytes,2,opt,name=HouseName,proto3" json:"HouseName,omitempty"`
	City        string `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNumber int64  `protobuf:"varint,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Error       string `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetUserResponce) Reset() {
	*x = GetUserResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponce) ProtoMessage() {}

func (x *GetUserResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponce.ProtoReflect.Descriptor instead.
func (*GetUserResponce) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserResponce) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserResponce) GetHouseName() string {
	if x != nil {
		return x.HouseName
	}
	return ""
}

func (x *GetUserResponce) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *GetUserResponce) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserResponce) GetPhoneNumber() int64 {
	if x != nil {
		return x.PhoneNumber
	}
	return 0
}

func (x *GetUserResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetAllUserDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUserDataReq) Reset() {
	*x = GetAllUserDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserDataReq) ProtoMessage() {}

func (x *GetAllUserDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserDataReq.ProtoReflect.Descriptor instead.
func (*GetAllUserDataReq) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{6}
}

type GetAllUserDataResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []string `protobuf:"bytes,1,rep,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetAllUserDataResult) Reset() {
	*x = GetAllUserDataResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_userproto_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserDataResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserDataResult) ProtoMessage() {}

func (x *GetAllUserDataResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_userproto_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserDataResult.ProtoReflect.Descriptor instead.
func (*GetAllUserDataResult) Descriptor() ([]byte, []int) {
	return file_pkg_pb_userproto_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllUserDataResult) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

var File_pkg_pb_userproto_user_proto protoreflect.FileDescriptor

var file_pkg_pb_userproto_user_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x13,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0xfb, 0x02, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x76, 0x63, 0x12, 0x46, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x14, 0x5a,
	0x12, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_userproto_user_proto_rawDescOnce sync.Once
	file_pkg_pb_userproto_user_proto_rawDescData = file_pkg_pb_userproto_user_proto_rawDesc
)

func file_pkg_pb_userproto_user_proto_rawDescGZIP() []byte {
	file_pkg_pb_userproto_user_proto_rawDescOnce.Do(func() {
		file_pkg_pb_userproto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_userproto_user_proto_rawDescData)
	})
	return file_pkg_pb_userproto_user_proto_rawDescData
}

var file_pkg_pb_userproto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_pb_userproto_user_proto_goTypes = []interface{}{
	(*CreatUserRequest)(nil),     // 0: userproto.CreatUserRequest
	(*UpdateUserRequest)(nil),    // 1: userproto.UpdateUserRequest
	(*GetUserRequest)(nil),       // 2: userproto.GetUserRequest
	(*DeleteUserRequest)(nil),    // 3: userproto.DeleteUserRequest
	(*CommonResponse)(nil),       // 4: userproto.CommonResponse
	(*GetUserResponce)(nil),      // 5: userproto.GetUserResponce
	(*GetAllUserDataReq)(nil),    // 6: userproto.GetAllUserDataReq
	(*GetAllUserDataResult)(nil), // 7: userproto.GetAllUserDataResult
}
var file_pkg_pb_userproto_user_proto_depIdxs = []int32{
	0, // 0: userproto.UserSvc.CreateUser:input_type -> userproto.CreatUserRequest
	1, // 1: userproto.UserSvc.UpdateUser:input_type -> userproto.UpdateUserRequest
	2, // 2: userproto.UserSvc.GetUserById:input_type -> userproto.GetUserRequest
	3, // 3: userproto.UserSvc.DeleteUser:input_type -> userproto.DeleteUserRequest
	6, // 4: userproto.UserSvc.GetUserData:input_type -> userproto.GetAllUserDataReq
	4, // 5: userproto.UserSvc.CreateUser:output_type -> userproto.CommonResponse
	4, // 6: userproto.UserSvc.UpdateUser:output_type -> userproto.CommonResponse
	5, // 7: userproto.UserSvc.GetUserById:output_type -> userproto.GetUserResponce
	4, // 8: userproto.UserSvc.DeleteUser:output_type -> userproto.CommonResponse
	7, // 9: userproto.UserSvc.GetUserData:output_type -> userproto.GetAllUserDataResult
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_userproto_user_proto_init() }
func file_pkg_pb_userproto_user_proto_init() {
	if File_pkg_pb_userproto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_userproto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatUserRequest); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponce); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserDataReq); i {
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
		file_pkg_pb_userproto_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserDataResult); i {
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
			RawDescriptor: file_pkg_pb_userproto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_userproto_user_proto_goTypes,
		DependencyIndexes: file_pkg_pb_userproto_user_proto_depIdxs,
		MessageInfos:      file_pkg_pb_userproto_user_proto_msgTypes,
	}.Build()
	File_pkg_pb_userproto_user_proto = out.File
	file_pkg_pb_userproto_user_proto_rawDesc = nil
	file_pkg_pb_userproto_user_proto_goTypes = nil
	file_pkg_pb_userproto_user_proto_depIdxs = nil
}
