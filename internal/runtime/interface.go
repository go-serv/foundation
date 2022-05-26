package runtime

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
)

type RuntimeInterface interface {
	NetworkServices() []i.NetworkServiceInterface
	MethodDescriptorByMessage(proto.Message) (i.MethodDescriptorInterface, error)
}
