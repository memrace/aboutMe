// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/api.proto

package service

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *UserReply) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserReply) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetDialog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetDialog) Reset() {
	*x = GetDialog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialog) ProtoMessage() {}

func (x *GetDialog) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialog.ProtoReflect.Descriptor instead.
func (*GetDialog) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetDialog) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Dialog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	ChatId    int64  `protobuf:"varint,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Reply     string `protobuf:"bytes,6,opt,name=reply,proto3" json:"reply,omitempty"`
	Replied   bool   `protobuf:"varint,7,opt,name=replied,proto3" json:"replied,omitempty"`
}

func (x *Dialog) Reset() {
	*x = Dialog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dialog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dialog) ProtoMessage() {}

func (x *Dialog) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dialog.ProtoReflect.Descriptor instead.
func (*Dialog) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *Dialog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dialog) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Dialog) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Dialog) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Dialog) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *Dialog) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

func (x *Dialog) GetReplied() bool {
	if x != nil {
		return x.Replied
	}
	return false
}

type CreateDialog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	ChatId    int64  `protobuf:"varint,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *CreateDialog) Reset() {
	*x = CreateDialog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDialog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDialog) ProtoMessage() {}

func (x *CreateDialog) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDialog.ProtoReflect.Descriptor instead.
func (*CreateDialog) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDialog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateDialog) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CreateDialog) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateDialog) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateDialog) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type DialogId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DialogId) Reset() {
	*x = DialogId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DialogId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DialogId) ProtoMessage() {}

func (x *DialogId) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DialogId.ProtoReflect.Descriptor instead.
func (*DialogId) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *DialogId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x24, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x06,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd6, 0x01, 0x0a, 0x0d, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x00, 0x12, 0x2e, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x08, 0x53, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x0f, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x40, 0x5a, 0x3e, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_api_proto_goTypes = []interface{}{
	(*Result)(nil),       // 0: service.Result
	(*UserReply)(nil),    // 1: service.UserReply
	(*GetDialog)(nil),    // 2: service.GetDialog
	(*Dialog)(nil),       // 3: service.Dialog
	(*CreateDialog)(nil), // 4: service.CreateDialog
	(*DialogId)(nil),     // 5: service.DialogId
}
var file_api_api_proto_depIdxs = []int32{
	2, // 0: service.DialogService.Get:input_type -> service.GetDialog
	4, // 1: service.DialogService.Create:input_type -> service.CreateDialog
	5, // 2: service.DialogService.Delete:input_type -> service.DialogId
	1, // 3: service.DialogService.SetReply:input_type -> service.UserReply
	3, // 4: service.DialogService.Get:output_type -> service.Dialog
	5, // 5: service.DialogService.Create:output_type -> service.DialogId
	0, // 6: service.DialogService.Delete:output_type -> service.Result
	0, // 7: service.DialogService.SetReply:output_type -> service.Result
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReply); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialog); i {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dialog); i {
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
		file_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDialog); i {
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
		file_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DialogId); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
