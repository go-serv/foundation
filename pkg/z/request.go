package z

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type MetaInterface interface {
	Dictionary() interface{}
	Hydrate() error
	Dehydrate() (metadata.MD, error)
	Copy(metaInterface MetaInterface)
}

type RequestResponseInterface interface {
	Data() any
	WithData(any)
	Meta() MetaInterface
	ServiceReflection() ServiceReflectionInterface
	MethodReflection() MethodReflectionInterface
	MessageReflection() MessageReflectionInterface
	Populate(proto.Message) error
}

type RequestInterface interface {
	RequestResponseInterface
}

type ResponseInterface interface {
	RequestResponseInterface
}

type ContextInterface interface {
	Interface() context.Context
	WithInterface(context.Context)
	Request() RequestInterface
	WithRequest(RequestInterface)
	Response() ResponseInterface
	WithResponse(ResponseInterface)
	Invoke() error
}

type NetContextInterface interface {
	ContextInterface
	NetworkService() NetworkServiceInterface
	Tenant() TenantId
	WithTenant(TenantId)
}

type NetServerContextInterface interface {
	NetContextInterface
	Session() SessionInterface
	WithSession(SessionInterface)
	Server() NetworkServerInterface
	WithServer(NetworkServerInterface)
}

type NetClientContextInterface interface {
	NetContextInterface
	WithClientInvoker(grpc.UnaryInvoker, *grpc.ClientConn, []grpc.CallOption)
	Client() NetworkClientInterface
	WithClient(NetworkClientInterface)
	AddCallOption(grpc.CallOption)
}

type LocalContextInterface interface {
	ContextInterface
}
