package z

import (
	"github.com/go-serv/service/pkg/z/platform"
	"google.golang.org/protobuf/proto"
)

type PlatformInterface interface {
	platform.FilesystemInterface
}

type RuntimeInterface interface {
	Platform() PlatformInterface
	RegisteredServices() []ServiceInterface
	RegisterLocalClient(LocalClientInterface)
	RegisterNetworkClient(NetworkClientInterface)
	Reflection() ReflectInterface
	NetworkServices() []NetworkServiceInterface
	IsRequestMessage(msg proto.Message) (bool, error)
	IsResponseMessage(msg proto.Message) (bool, error)

	RegisterEventHandler(func(eventTyp interface{}))
	TriggerEvent(eventTyp interface{}, extra ...interface{})
}
