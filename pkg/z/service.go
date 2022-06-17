package z

import (
	"crypto/md5"
	"encoding/binary"
	"github.com/go-serv/service/pkg/z/platform"
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
	EncriptionKey() []byte
	WithEncriptionKey([]byte)
}

type LocalServiceInterface interface {
	ServiceInterface
}

type FtpUploadProfileInterface interface {
	RootDir() platform.Pathname
	MaxFileSize() int64
	FilePerms() uint32
}
