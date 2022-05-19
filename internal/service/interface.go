package service

import (
	"github.com/go-serv/service/internal/grpc/descriptor"
	"github.com/go-serv/service/internal/grpc/request"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ConfigInterface interface {
}

type GrpcMessageWrapperFn func(in []byte) []byte

type baseServiceInterface interface {
	Service_Descriptor() descriptor.ServiceDescriptorInterface
	Service_AddServiceProtoExtension(info *protoimpl.ExtensionInfo)
	Service_AddMethodProtoExtension(info *protoimpl.ExtensionInfo)
	Service_Register(srv *grpc.Server)
}

type LocalServiceInterface interface {
	baseServiceInterface
}

type NetworkServiceInterface interface {
	baseServiceInterface
	Service_OnNewSession(req request.RequestInterface) error
	Service_RequestNewSession(req request.RequestInterface) int32
}
