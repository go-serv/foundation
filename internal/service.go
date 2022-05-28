package internal

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ServiceInterface interface {
	Name() protoreflect.FullName
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
	CodecMiddlewareGroup() CodecMiddlewareGroupInterface
	Service_Descriptor() ServiceDescriptorInterface
	Service_AddServiceProtoExtension(info *protoimpl.ExtensionInfo)
	Service_AddMethodProtoExtension(info *protoimpl.ExtensionInfo)
	Service_Register(srv *grpc.Server)
}

type NetworkServiceInterface interface {
	ServiceInterface
	Service_OnNewSession(req RequestInterface) error
	// Service_InfoNewSession returns timeout in seconds for a new session. Zero means no new session is required
	Service_InfoNewSession(methodName string) int32
	Service_InfoMsgEncryption(methodName string) bool

	Service_EncriptionKey() []byte
	Service_WithEncriptionKey([]byte)
}

type LocalServiceInterface interface {
	ServiceInterface
}
