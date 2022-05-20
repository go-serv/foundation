package service

import (
	"github.com/go-serv/service/internal/grpc/descriptor"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ConfigInterface interface {
}

type BaseServiceInterface interface {
	Service_Descriptor() descriptor.ServiceDescriptorInterface
	Service_AddServiceProtoExtension(info *protoimpl.ExtensionInfo)
	Service_AddMethodProtoExtension(info *protoimpl.ExtensionInfo)
	Service_Register(srv *grpc.Server)
}
