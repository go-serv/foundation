package z

import (
	"crypto/md5"
	"encoding/binary"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func BaseErrorCodeFromServiceName(svcName string) uint64 {
	hash := md5.Sum([]byte(svcName))
	return binary.LittleEndian.Uint64(hash[0:8])
}

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
