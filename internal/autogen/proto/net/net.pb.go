// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: net/net.proto

package net

import (
	_ "github.com/go-serv/foundation/internal/autogen/proto/go_serv"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_net_net_proto protoreflect.FileDescriptor

var file_net_net_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x1a, 0x17, 0x67, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6e, 0x65, 0x74, 0x2f, 0x66, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb5, 0x03, 0x0a, 0x09, 0x4e, 0x65, 0x74, 0x50,
	0x61, 0x72, 0x63, 0x65, 0x6c, 0x12, 0x3f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e,
	0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0d, 0x46, 0x74, 0x70, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67,
	0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x4e,
	0x65, 0x77, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x04, 0xb0, 0xac, 0x1d, 0x01, 0x12, 0x5c, 0x0a, 0x0b, 0x46, 0x74, 0x70, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x04, 0xa8, 0xac, 0x1d, 0x01, 0x12, 0x57, 0x0a, 0x0a, 0x46, 0x74, 0x70, 0x49, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x46, 0x74, 0x70, 0x2e, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0xa8, 0xac, 0x1d, 0x01, 0x42,
	0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x2f,
	0x6e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_net_net_proto_goTypes = []interface{}{
	(*Ping_Request)(nil),            // 0: go_serv.net.Ping.Request
	(*Session_Request)(nil),         // 1: go_serv.net.Session.Request
	(*Ftp_NewSession_Request)(nil),  // 2: go_serv.net.Ftp.NewSession.Request
	(*Ftp_FileChunk_Request)(nil),   // 3: go_serv.net.Ftp.FileChunk.Request
	(*Ftp_Inquiry_Request)(nil),     // 4: go_serv.net.Ftp.Inquiry.Request
	(*Ping_Response)(nil),           // 5: go_serv.net.Ping.Response
	(*Session_Response)(nil),        // 6: go_serv.net.Session.Response
	(*Ftp_NewSession_Response)(nil), // 7: go_serv.net.Ftp.NewSession.Response
	(*Ftp_FileChunk_Response)(nil),  // 8: go_serv.net.Ftp.FileChunk.Response
	(*Ftp_Inquiry_Response)(nil),    // 9: go_serv.net.Ftp.Inquiry.Response
}
var file_net_net_proto_depIdxs = []int32{
	0, // 0: go_serv.net.NetParcel.Ping:input_type -> go_serv.net.Ping.Request
	1, // 1: go_serv.net.NetParcel.SecureSession:input_type -> go_serv.net.Session.Request
	2, // 2: go_serv.net.NetParcel.FtpNewSession:input_type -> go_serv.net.Ftp.NewSession.Request
	3, // 3: go_serv.net.NetParcel.FtpTransfer:input_type -> go_serv.net.Ftp.FileChunk.Request
	4, // 4: go_serv.net.NetParcel.FtpInquiry:input_type -> go_serv.net.Ftp.Inquiry.Request
	5, // 5: go_serv.net.NetParcel.Ping:output_type -> go_serv.net.Ping.Response
	6, // 6: go_serv.net.NetParcel.SecureSession:output_type -> go_serv.net.Session.Response
	7, // 7: go_serv.net.NetParcel.FtpNewSession:output_type -> go_serv.net.Ftp.NewSession.Response
	8, // 8: go_serv.net.NetParcel.FtpTransfer:output_type -> go_serv.net.Ftp.FileChunk.Response
	9, // 9: go_serv.net.NetParcel.FtpInquiry:output_type -> go_serv.net.Ftp.Inquiry.Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_net_net_proto_init() }
func file_net_net_proto_init() {
	if File_net_net_proto != nil {
		return
	}
	file_net_ping_proto_init()
	file_net_ftp_proto_init()
	file_net_session_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_net_net_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_net_net_proto_goTypes,
		DependencyIndexes: file_net_net_proto_depIdxs,
	}.Build()
	File_net_net_proto = out.File
	file_net_net_proto_rawDesc = nil
	file_net_net_proto_goTypes = nil
	file_net_net_proto_depIdxs = nil
}
