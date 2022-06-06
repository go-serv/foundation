package z

import "google.golang.org/protobuf/proto"

type RuntimeInterface interface {
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
