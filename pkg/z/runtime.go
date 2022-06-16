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

type ResolverInterface interface {
	Run(args ...any) (v any, err error)
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
	// AddResolver adds a value resolver by the given key.
	AddResolver(key any, resolver ResolverInterface)
	// Resolve executes the resolver handler only once and returns a value returned by it.
	Resolve(key any, args ...any) (any, error)
}
