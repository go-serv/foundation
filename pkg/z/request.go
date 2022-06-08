package z

import "google.golang.org/grpc/metadata"

type MetaInterface interface {
	Dictionary() interface{}
	Hydrate() error
	Dehydrate() (metadata.MD, error)
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
	Response() ResponseInterface
	Invoke() (interface{}, error)
}

type NetContextInterface interface {
	ContextInterface
	Session() SessionInterface
}

type LocalContextInterface interface {
	ContextInterface
}
