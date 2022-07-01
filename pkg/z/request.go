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
	DataFrame() DataFrameInterface
	Meta() MetaInterface
	MethodReflection() MethodReflectionInterface
	MessageReflection() MessageReflectionInterface
}

type RequestInterface interface {
	RequestResponseInterface
}

type ResponseInterface interface {
	RequestResponseInterface
	WithDataFrame(DataFrameInterface)
	Populate(proto.Message) error
}

type ContextInterface interface {
	WithInput(proto.Message) error
	WithOutput(proto.Message) error
	MethodReflection() MethodReflectionInterface
	InputReflection() MessageReflectionInterface
	OutputReflection() MessageReflectionInterface
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
