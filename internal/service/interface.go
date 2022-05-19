package service

import (
	"github.com/go-serv/service/internal/grpc/request"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ConfigInterface interface {
}

type GrpcMessageWrapperFn func(in []byte) []byte

type baseServiceInterface interface {
	Service_AddServiceProtoExtension(info *protoimpl.ExtensionInfo, defaultVal interface{})
	Service_AddMethodProtoExtension(info *protoimpl.ExtensionInfo, defaultVal interface{})
	Service_ServiceProtoExtensions() []*serviceProtoExt
	Service_MethodProtoExtensions() []*methodProtoExt
	Service_Register(srv *grpc.Server)
}

type ServiceReflectionInterface interface {
	AddServiceProtoExtension(info *protoimpl.ExtensionInfo, defaultVal interface{})
	AddMethodProtoExtension(info *protoimpl.ExtensionInfo, defaultVal interface{})
	ServiceProtoExtensions() []*serviceProtoExt
	MethodProtoExtensions() []*methodProtoExt
}

type LocalServiceInterface interface {
	baseServiceInterface
}

type NetworkServiceInterface interface {
	baseServiceInterface
	Service_OnNewSession(req request.RequestInterface) error
}
