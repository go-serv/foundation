package internal

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ServiceInterface interface {
	Name() protoreflect.FullName
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
	CodecMiddlewareGroup() CodecMiddlewareGroupInterface
	Register(srv *grpc.Server)
}

type NetworkServiceInterface interface {
	ServiceInterface
	Service_OnNewSession(req RequestInterface) error
	// Service_InfoNewSession returns timeout in seconds for a new session. Zero means no new session is required
	Service_InfoNewSession(methodName string) int32
	Service_InfoMsgEncryption(methodName string) bool
	EncriptionKey() []byte
	WithEncriptionKey([]byte)
}

type LocalServiceInterface interface {
	ServiceInterface
}
