package z

import (
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
	WithData(any)
}

type ContextInterface interface {
	Request() RequestInterface
	WithRequest(RequestInterface)
	Response() ResponseInterface
	WithResponse(ResponseInterface)
	Invoke() error
}

type NetContextInterface interface {
	ContextInterface
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
}

type LocalContextInterface interface {
	ContextInterface
}
