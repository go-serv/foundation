package z

import (
	"google.golang.org/protobuf/proto"
)

type (
	FtpUploadProfilesResolverTyp int
)

const (
	FtpUploadProfilerResolver FtpUploadProfilesResolverTyp = iota + 1
)

type RuntimeInterface interface {
	RegisterService(ServiceInterface)
	RegisterClient(ClientInterface)

	Services() []ServiceInterface

	Platform() PlatformInterface

	RegisterNetworkClient(NetworkClientInterface)

	Reflection() ReflectInterface
	NetworkServices() []NetworkServiceInterface
	IsRequestMessage(msg proto.Message) (bool, error)
	IsResponseMessage(msg proto.Message) (bool, error)
	RegisterEventHandler(func(eventTyp interface{}))
	TriggerEvent(eventTyp interface{}, extra ...interface{})
	// AddResolver adds a value resolver with the given key.
	AddResolver(key any, resolver MemoizerInterface)
	// Resolve executes the resolver handler only once and returns a value returned by it.
	Resolve(key any, args ...any) (any, error)
}
