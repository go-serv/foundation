package z

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type MetaInterface interface {
	Dictionary() interface{}
	Hydrate() error
	Dehydrate() (metadata.MD, error)
	Copy(metaInterface MetaInterface)
}

type RequestResponseInterface interface {
	Payload() interface{}
	WithPayload(interface{})
	Meta() MetaInterface
	MethodReflection() MethodReflectionInterface
	MessageReflection() MessageReflectionInterface
}

type RequestInterface interface {
	RequestResponseInterface
}

type ResponseInterface interface {
	RequestResponseInterface
	ToGrpcResponse() interface{}
}

type ContextInterface interface {
	Request() RequestInterface
	WithRequest(RequestInterface)
	Response() ResponseInterface
	WithResponse(ResponseInterface)
	Invoke() (interface{}, error)
}

type NetContextInterface interface {
	ContextInterface
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
}

type LocalContextInterface interface {
	ContextInterface
}
