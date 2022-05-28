package runtime

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type RuntimeInterface interface {
	NetworkServices() []i.NetworkServiceInterface
	MethodDescriptorByMessage(proto.Message) (i.MethodDescriptorInterface, error)
	IsRequestMessage(msg proto.Message) (bool, error)
	IsResponseMessage(msg proto.Message) (bool, error)

	RegisterLocalClient(svcName protoreflect.FullName, c i.LocalClientInterface)
	RegisterNetworkClient(svcName protoreflect.FullName, c i.NetworkClientInterface)
}
