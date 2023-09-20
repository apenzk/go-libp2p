// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: msg.proto

package pb

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

// Specifies type in `data` field.
type Message_Type int32

const (
	// String of `RTCSessionDescription.sdp`
	Message_SDP_OFFER Message_Type = 0
	// String of `RTCSessionDescription.sdp`
	Message_SDP_ANSWER Message_Type = 1
	// String of `RTCIceCandidate.toJSON()`
	Message_ICE_CANDIDATE Message_Type = 2
)

// Enum value maps for Message_Type.
var (
	Message_Type_name = map[int32]string{
		0: "SDP_OFFER",
		1: "SDP_ANSWER",
		2: "ICE_CANDIDATE",
	}
	Message_Type_value = map[string]int32{
		"SDP_OFFER":     0,
		"SDP_ANSWER":    1,
		"ICE_CANDIDATE": 2,
	}
)

func (x Message_Type) Enum() *Message_Type {
	p := new(Message_Type)
	*p = x
	return p
}

func (x Message_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (Message_Type) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x Message_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Message_Type.Descriptor instead.
func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0, 0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *Message_Type `protobuf:"varint,1,opt,name=type,proto3,enum=libp2pwebrtcprivate.pb.Message_Type,oneof" json:"type,omitempty"`
	Data *string       `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() Message_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Message_SDP_OFFER
}

func (x *Message) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6c, 0x69, 0x62,
	0x70, 0x32, 0x70, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x62, 0x22, 0xad, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x22, 0x38, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x44, 0x50, 0x5f, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x44, 0x50, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x44, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x70,
	0x32, 0x70, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x77, 0x65, 0x62, 0x72, 0x74, 0x63, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_msg_proto_goTypes = []interface{}{
	(Message_Type)(0), // 0: libp2pwebrtcprivate.pb.Message.Type
	(*Message)(nil),   // 1: libp2pwebrtcprivate.pb.Message
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: libp2pwebrtcprivate.pb.Message.type:type_name -> libp2pwebrtcprivate.pb.Message.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	file_msg_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
