package internal

import (
	"google.golang.org/grpc"
)

type MiddlewareGroupInterface interface {
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	UnaryClientInterceptor() grpc.UnaryClientInterceptor
}

type (
	NetPreStreamHandlerFn func(ServiceReflectInterface, MethodReflectionInterface, MessageReflectionInterface) error
	NetRequestHandlerFn   func(next NetChainElement, req RequestInterface, res ResponseInterface) error
	NetResponseHandlerFn  func(next NetChainElement, res ResponseInterface) error
	NetChainElement       func(RequestInterface, ResponseInterface) (NetChainElement, error)
)
type NetMiddlewareGroupInterface interface {
	AddPreStreamHandler(fn NetPreStreamHandlerFn)
	AddRequestHandler(fn NetRequestHandlerFn)
	AddResponseHandler(fn NetResponseHandlerFn)
}

type MetaInterface interface {
	Dictionary() interface{}
	Hydrate() error
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

type RequestDataInterface interface {
	Validate() bool
	Errors() []error
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

type SessionInterface interface {
	Id() SessionId
}
