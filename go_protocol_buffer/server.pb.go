// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/server.proto

package go_protocol_buffer

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

type ShareGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LinkKey     string  `protobuf:"bytes,2,opt,name=linkKey,proto3" json:"linkKey,omitempty"`
	DestLon     float64 `protobuf:"fixed64,3,opt,name=destLon,proto3" json:"destLon,omitempty"`
	DestLat     float64 `protobuf:"fixed64,4,opt,name=destLat,proto3" json:"destLat,omitempty"`
	Users       []*User `protobuf:"bytes,5,rep,name=users,proto3" json:"users,omitempty"`
	MeetingTime string  `protobuf:"bytes,6,opt,name=meetingTime,proto3" json:"meetingTime,omitempty"`
	InviteUrl   string  `protobuf:"bytes,7,opt,name=inviteUrl,proto3" json:"inviteUrl,omitempty"`
}

func (x *ShareGroup) Reset() {
	*x = ShareGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareGroup) ProtoMessage() {}

func (x *ShareGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareGroup.ProtoReflect.Descriptor instead.
func (*ShareGroup) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *ShareGroup) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ShareGroup) GetLinkKey() string {
	if x != nil {
		return x.LinkKey
	}
	return ""
}

func (x *ShareGroup) GetDestLon() float64 {
	if x != nil {
		return x.DestLon
	}
	return 0
}

func (x *ShareGroup) GetDestLat() float64 {
	if x != nil {
		return x.DestLat
	}
	return 0
}

func (x *ShareGroup) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ShareGroup) GetMeetingTime() string {
	if x != nil {
		return x.MeetingTime
	}
	return ""
}

func (x *ShareGroup) GetInviteUrl() string {
	if x != nil {
		return x.InviteUrl
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShareGroup   *ShareGroup `protobuf:"bytes,3,opt,name=shareGroup,proto3" json:"shareGroup,omitempty"`
	ShareGroupId uint64      `protobuf:"varint,4,opt,name=shareGroupId,proto3" json:"shareGroupId,omitempty"`
	Lat          float64     `protobuf:"fixed64,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon          float64     `protobuf:"fixed64,6,opt,name=lon,proto3" json:"lon,omitempty"`
	PositionAt   string      `protobuf:"bytes,7,opt,name=positionAt,proto3" json:"positionAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetShareGroup() *ShareGroup {
	if x != nil {
		return x.ShareGroup
	}
	return nil
}

func (x *User) GetShareGroupId() uint64 {
	if x != nil {
		return x.ShareGroupId
	}
	return 0
}

func (x *User) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *User) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *User) GetPositionAt() string {
	if x != nil {
		return x.PositionAt
	}
	return ""
}

type AnonymousSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AnonymousSignUpRequest) Reset() {
	*x = AnonymousSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnonymousSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnonymousSignUpRequest) ProtoMessage() {}

func (x *AnonymousSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnonymousSignUpRequest.ProtoReflect.Descriptor instead.
func (*AnonymousSignUpRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{2}
}

func (x *AnonymousSignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AnonymousSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	AccessToken string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
}

func (x *AnonymousSignUpResponse) Reset() {
	*x = AnonymousSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnonymousSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnonymousSignUpResponse) ProtoMessage() {}

func (x *AnonymousSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnonymousSignUpResponse.ProtoReflect.Descriptor instead.
func (*AnonymousSignUpResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{3}
}

func (x *AnonymousSignUpResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AnonymousSignUpResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type CreateShareGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestLon     float64 `protobuf:"fixed64,3,opt,name=destLon,proto3" json:"destLon,omitempty"`
	DestLat     float64 `protobuf:"fixed64,4,opt,name=destLat,proto3" json:"destLat,omitempty"`
	MeetingTime string  `protobuf:"bytes,5,opt,name=meetingTime,proto3" json:"meetingTime,omitempty"`
}

func (x *CreateShareGroupRequest) Reset() {
	*x = CreateShareGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShareGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShareGroupRequest) ProtoMessage() {}

func (x *CreateShareGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShareGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateShareGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{4}
}

func (x *CreateShareGroupRequest) GetDestLon() float64 {
	if x != nil {
		return x.DestLon
	}
	return 0
}

func (x *CreateShareGroupRequest) GetDestLat() float64 {
	if x != nil {
		return x.DestLat
	}
	return 0
}

func (x *CreateShareGroupRequest) GetMeetingTime() string {
	if x != nil {
		return x.MeetingTime
	}
	return ""
}

type CreateShareGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareGroup *ShareGroup `protobuf:"bytes,1,opt,name=shareGroup,proto3" json:"shareGroup,omitempty"`
}

func (x *CreateShareGroupResponse) Reset() {
	*x = CreateShareGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShareGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShareGroupResponse) ProtoMessage() {}

func (x *CreateShareGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShareGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateShareGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{5}
}

func (x *CreateShareGroupResponse) GetShareGroup() *ShareGroup {
	if x != nil {
		return x.ShareGroup
	}
	return nil
}

type JoinShareGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkKey string `protobuf:"bytes,1,opt,name=linkKey,proto3" json:"linkKey,omitempty"`
}

func (x *JoinShareGroupRequest) Reset() {
	*x = JoinShareGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinShareGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinShareGroupRequest) ProtoMessage() {}

func (x *JoinShareGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinShareGroupRequest.ProtoReflect.Descriptor instead.
func (*JoinShareGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{6}
}

func (x *JoinShareGroupRequest) GetLinkKey() string {
	if x != nil {
		return x.LinkKey
	}
	return ""
}

type JoinShareGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareGroup *ShareGroup `protobuf:"bytes,1,opt,name=shareGroup,proto3" json:"shareGroup,omitempty"`
}

func (x *JoinShareGroupResponse) Reset() {
	*x = JoinShareGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinShareGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinShareGroupResponse) ProtoMessage() {}

func (x *JoinShareGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinShareGroupResponse.ProtoReflect.Descriptor instead.
func (*JoinShareGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{7}
}

func (x *JoinShareGroupResponse) GetShareGroup() *ShareGroup {
	if x != nil {
		return x.ShareGroup
	}
	return nil
}

type GetCurrentShareGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *GetCurrentShareGroupRequest) Reset() {
	*x = GetCurrentShareGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentShareGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentShareGroupRequest) ProtoMessage() {}

func (x *GetCurrentShareGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentShareGroupRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentShareGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{8}
}

func (x *GetCurrentShareGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetCurrentShareGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareGroup *ShareGroup `protobuf:"bytes,1,opt,name=shareGroup,proto3" json:"shareGroup,omitempty"`
}

func (x *GetCurrentShareGroupResponse) Reset() {
	*x = GetCurrentShareGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentShareGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentShareGroupResponse) ProtoMessage() {}

func (x *GetCurrentShareGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentShareGroupResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentShareGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{9}
}

func (x *GetCurrentShareGroupResponse) GetShareGroup() *ShareGroup {
	if x != nil {
		return x.ShareGroup
	}
	return nil
}

type UpdatePositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
}

func (x *UpdatePositionRequest) Reset() {
	*x = UpdatePositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePositionRequest) ProtoMessage() {}

func (x *UpdatePositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePositionRequest.ProtoReflect.Descriptor instead.
func (*UpdatePositionRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{10}
}

func (x *UpdatePositionRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *UpdatePositionRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type UpdatePositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdatePositionResponse) Reset() {
	*x = UpdatePositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePositionResponse) ProtoMessage() {}

func (x *UpdatePositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePositionResponse.ProtoReflect.Descriptor instead.
func (*UpdatePositionResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{11}
}

func (x *UpdatePositionResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetCurrentUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentUserRequest) Reset() {
	*x = GetCurrentUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentUserRequest) ProtoMessage() {}

func (x *GetCurrentUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentUserRequest.ProtoReflect.Descriptor instead.
func (*GetCurrentUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{12}
}

type GetCurrentUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetCurrentUserResponse) Reset() {
	*x = GetCurrentUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentUserResponse) ProtoMessage() {}

func (x *GetCurrentUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentUserResponse.ProtoReflect.Descriptor instead.
func (*GetCurrentUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{13}
}

func (x *GetCurrentUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0xce, 0x01, 0x0a,
	0x0a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c,
	0x69, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69,
	0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xc6, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x16, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x17, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x6f, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x64, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74,
	0x4c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x4c,
	0x61, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x31, 0x0a, 0x15, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6c, 0x69, 0x6e, 0x6b, 0x4b, 0x65, 0x79, 0x22, 0x4c, 0x0a, 0x16, 0x4a, 0x6f, 0x69, 0x6e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x37, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x52,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x3b, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x22,
	0x3a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x32, 0x8a, 0x04, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0f,
	0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12,
	0x1e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x4a, 0x6f, 0x69, 0x6e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1d, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x23, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a,
	0x14, 0x2e, 0x2f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_server_proto_goTypes = []interface{}{
	(*ShareGroup)(nil),                   // 0: Server.ShareGroup
	(*User)(nil),                         // 1: Server.User
	(*AnonymousSignUpRequest)(nil),       // 2: Server.AnonymousSignUpRequest
	(*AnonymousSignUpResponse)(nil),      // 3: Server.AnonymousSignUpResponse
	(*CreateShareGroupRequest)(nil),      // 4: Server.CreateShareGroupRequest
	(*CreateShareGroupResponse)(nil),     // 5: Server.CreateShareGroupResponse
	(*JoinShareGroupRequest)(nil),        // 6: Server.JoinShareGroupRequest
	(*JoinShareGroupResponse)(nil),       // 7: Server.JoinShareGroupResponse
	(*GetCurrentShareGroupRequest)(nil),  // 8: Server.GetCurrentShareGroupRequest
	(*GetCurrentShareGroupResponse)(nil), // 9: Server.GetCurrentShareGroupResponse
	(*UpdatePositionRequest)(nil),        // 10: Server.UpdatePositionRequest
	(*UpdatePositionResponse)(nil),       // 11: Server.UpdatePositionResponse
	(*GetCurrentUserRequest)(nil),        // 12: Server.GetCurrentUserRequest
	(*GetCurrentUserResponse)(nil),       // 13: Server.GetCurrentUserResponse
}
var file_proto_server_proto_depIdxs = []int32{
	1,  // 0: Server.ShareGroup.users:type_name -> Server.User
	0,  // 1: Server.User.shareGroup:type_name -> Server.ShareGroup
	1,  // 2: Server.AnonymousSignUpResponse.user:type_name -> Server.User
	0,  // 3: Server.CreateShareGroupResponse.shareGroup:type_name -> Server.ShareGroup
	0,  // 4: Server.JoinShareGroupResponse.shareGroup:type_name -> Server.ShareGroup
	0,  // 5: Server.GetCurrentShareGroupResponse.shareGroup:type_name -> Server.ShareGroup
	1,  // 6: Server.UpdatePositionResponse.user:type_name -> Server.User
	1,  // 7: Server.GetCurrentUserResponse.user:type_name -> Server.User
	2,  // 8: Server.Service.AnonymousSignUp:input_type -> Server.AnonymousSignUpRequest
	4,  // 9: Server.Service.CreateShareGroup:input_type -> Server.CreateShareGroupRequest
	6,  // 10: Server.Service.JoinShareGroup:input_type -> Server.JoinShareGroupRequest
	8,  // 11: Server.Service.GetCurrentShareGroup:input_type -> Server.GetCurrentShareGroupRequest
	10, // 12: Server.Service.UpdatePosition:input_type -> Server.UpdatePositionRequest
	12, // 13: Server.Service.GetCurrentUser:input_type -> Server.GetCurrentUserRequest
	3,  // 14: Server.Service.AnonymousSignUp:output_type -> Server.AnonymousSignUpResponse
	5,  // 15: Server.Service.CreateShareGroup:output_type -> Server.CreateShareGroupResponse
	7,  // 16: Server.Service.JoinShareGroup:output_type -> Server.JoinShareGroupResponse
	9,  // 17: Server.Service.GetCurrentShareGroup:output_type -> Server.GetCurrentShareGroupResponse
	11, // 18: Server.Service.UpdatePosition:output_type -> Server.UpdatePositionResponse
	13, // 19: Server.Service.GetCurrentUser:output_type -> Server.GetCurrentUserResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareGroup); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnonymousSignUpRequest); i {
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
		file_proto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnonymousSignUpResponse); i {
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
		file_proto_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShareGroupRequest); i {
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
		file_proto_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShareGroupResponse); i {
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
		file_proto_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinShareGroupRequest); i {
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
		file_proto_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinShareGroupResponse); i {
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
		file_proto_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentShareGroupRequest); i {
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
		file_proto_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentShareGroupResponse); i {
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
		file_proto_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePositionRequest); i {
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
		file_proto_server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePositionResponse); i {
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
		file_proto_server_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentUserRequest); i {
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
		file_proto_server_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentUserResponse); i {
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
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}
