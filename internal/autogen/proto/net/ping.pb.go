// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: net/ping.proto

package net

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_net_ping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_net_ping_proto_rawDescGZIP(), []int{0}
}

type Ping_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload uint64 `protobuf:"varint,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Ping_Request) Reset() {
	*x = Ping_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping_Request) ProtoMessage() {}

func (x *Ping_Request) ProtoReflect() protoreflect.Message {
	mi := &file_net_ping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping_Request.ProtoReflect.Descriptor instead.
func (*Ping_Request) Descriptor() ([]byte, []int) {
	return file_net_ping_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ping_Request) GetPayload() uint64 {
	if x != nil {
		return x.Payload
	}
	return 0
}

type Ping_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload uint64 `protobuf:"varint,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Ping_Response) Reset() {
	*x = Ping_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_ping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping_Response) ProtoMessage() {}

func (x *Ping_Response) ProtoReflect() protoreflect.Message {
	mi := &file_net_ping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping_Response.ProtoReflect.Descriptor instead.
func (*Ping_Response) Descriptor() ([]byte, []int) {
	return file_net_ping_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Ping_Response) GetPayload() uint64 {
	if x != nil {
		return x.Payload
	}
	return 0
}

var File_net_ping_proto protoreflect.FileDescriptor

var file_net_ping_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x22, 0x51, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x23, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x24, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x2f, 0x6e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_net_ping_proto_rawDescOnce sync.Once
	file_net_ping_proto_rawDescData = file_net_ping_proto_rawDesc
)

func file_net_ping_proto_rawDescGZIP() []byte {
	file_net_ping_proto_rawDescOnce.Do(func() {
		file_net_ping_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_ping_proto_rawDescData)
	})
	return file_net_ping_proto_rawDescData
}

var file_net_ping_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_net_ping_proto_goTypes = []interface{}{
	(*Ping)(nil),          // 0: go_serv.net.Ping
	(*Ping_Request)(nil),  // 1: go_serv.net.Ping.Request
	(*Ping_Response)(nil), // 2: go_serv.net.Ping.Response
}
var file_net_ping_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_net_ping_proto_init() }
func file_net_ping_proto_init() {
	if File_net_ping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_ping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_net_ping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping_Request); i {
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
		file_net_ping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping_Response); i {
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
			RawDescriptor: file_net_ping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_net_ping_proto_goTypes,
		DependencyIndexes: file_net_ping_proto_depIdxs,
		MessageInfos:      file_net_ping_proto_msgTypes,
	}.Build()
	File_net_ping_proto = out.File
	file_net_ping_proto_rawDesc = nil
	file_net_ping_proto_goTypes = nil
	file_net_ping_proto_depIdxs = nil
}
