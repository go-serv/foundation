package runtime

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
)

type RuntimeInterface interface {
	NetworkServices() []i.NetworkServiceInterface
	MethodDescriptorByMessage(proto.Message) (i.MethodDescriptorInterface, error)
	IsRequestMessage(msg proto.Message) (bool, error)
	IsResponseMessage(msg proto.Message) (bool, error)

	RegisterLocalClient(i.LocalClientInterface)
	RegisterNetworkClient(i.NetworkClientInterface)
}
