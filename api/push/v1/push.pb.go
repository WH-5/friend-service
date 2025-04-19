// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: api/push/v1/push.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushMsgRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ToUnique      string                 `protobuf:"bytes,1,opt,name=to_unique,json=toUnique,proto3" json:"to_unique,omitempty"`
	SelfUserId    uint64                 `protobuf:"varint,4,opt,name=self_user_id,json=selfUserId,proto3" json:"self_user_id,omitempty"`
	MsgType       int32                  `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"` //1聊天消息 2好友关系消息
	Payload       []byte                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`                 //消息内容
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushMsgRequest) Reset() {
	*x = PushMsgRequest{}
	mi := &file_api_push_v1_push_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgRequest) ProtoMessage() {}

func (x *PushMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_push_v1_push_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgRequest.ProtoReflect.Descriptor instead.
func (*PushMsgRequest) Descriptor() ([]byte, []int) {
	return file_api_push_v1_push_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgRequest) GetToUnique() string {
	if x != nil {
		return x.ToUnique
	}
	return ""
}

func (x *PushMsgRequest) GetSelfUserId() uint64 {
	if x != nil {
		return x.SelfUserId
	}
	return 0
}

func (x *PushMsgRequest) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *PushMsgRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PushMsgReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PushMsgReply) Reset() {
	*x = PushMsgReply{}
	mi := &file_api_push_v1_push_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PushMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReply) ProtoMessage() {}

func (x *PushMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_push_v1_push_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReply.ProtoReflect.Descriptor instead.
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return file_api_push_v1_push_proto_rawDescGZIP(), []int{1}
}

func (x *PushMsgReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ChatPayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 聊天消息
	Content       string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ContentType   string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatPayload) Reset() {
	*x = ChatPayload{}
	mi := &file_api_push_v1_push_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatPayload) ProtoMessage() {}

func (x *ChatPayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_push_v1_push_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatPayload.ProtoReflect.Descriptor instead.
func (*ChatPayload) Descriptor() ([]byte, []int) {
	return file_api_push_v1_push_proto_rawDescGZIP(), []int{2}
}

func (x *ChatPayload) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatPayload) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type FriendPayload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          int32                  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"` //1好友请求 2好友请求接受 3好友请求拒绝 4好友删除
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FriendPayload) Reset() {
	*x = FriendPayload{}
	mi := &file_api_push_v1_push_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendPayload) ProtoMessage() {}

func (x *FriendPayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_push_v1_push_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendPayload.ProtoReflect.Descriptor instead.
func (*FriendPayload) Descriptor() ([]byte, []int) {
	return file_api_push_v1_push_proto_rawDescGZIP(), []int{3}
}

func (x *FriendPayload) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_api_push_v1_push_proto protoreflect.FileDescriptor

var file_api_push_v1_push_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x6f,
	0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa,
	0x42, 0x1d, 0x72, 0x1b, 0x10, 0x01, 0x18, 0x14, 0x32, 0x15, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41,
	0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x7b, 0x31, 0x2c, 0x32, 0x30, 0x7d, 0x24, 0x52,
	0x08, 0x74, 0x6f, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x65, 0x6c,
	0x66, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x73, 0x65, 0x6c, 0x66, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d,
	0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x20, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x4a, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x23,
	0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x32, 0x49, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x41, 0x0a, 0x07, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3e,
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x48, 0x2d, 0x35,
	0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_push_v1_push_proto_rawDescOnce sync.Once
	file_api_push_v1_push_proto_rawDescData []byte
)

func file_api_push_v1_push_proto_rawDescGZIP() []byte {
	file_api_push_v1_push_proto_rawDescOnce.Do(func() {
		file_api_push_v1_push_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_push_v1_push_proto_rawDesc), len(file_api_push_v1_push_proto_rawDesc)))
	})
	return file_api_push_v1_push_proto_rawDescData
}

var file_api_push_v1_push_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_push_v1_push_proto_goTypes = []any{
	(*PushMsgRequest)(nil), // 0: api.push.v1.PushMsgRequest
	(*PushMsgReply)(nil),   // 1: api.push.v1.PushMsgReply
	(*ChatPayload)(nil),    // 2: api.push.v1.ChatPayload
	(*FriendPayload)(nil),  // 3: api.push.v1.FriendPayload
}
var file_api_push_v1_push_proto_depIdxs = []int32{
	0, // 0: api.push.v1.Push.PushMsg:input_type -> api.push.v1.PushMsgRequest
	1, // 1: api.push.v1.Push.PushMsg:output_type -> api.push.v1.PushMsgReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_push_v1_push_proto_init() }
func file_api_push_v1_push_proto_init() {
	if File_api_push_v1_push_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_push_v1_push_proto_rawDesc), len(file_api_push_v1_push_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_push_v1_push_proto_goTypes,
		DependencyIndexes: file_api_push_v1_push_proto_depIdxs,
		MessageInfos:      file_api_push_v1_push_proto_msgTypes,
	}.Build()
	File_api_push_v1_push_proto = out.File
	file_api_push_v1_push_proto_goTypes = nil
	file_api_push_v1_push_proto_depIdxs = nil
}
