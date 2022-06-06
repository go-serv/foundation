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
	NetResponseHandlerFn  func(next NetChainElement, res ResponseInterface) (interface{}, error)
	NetChainElement       func(RequestInterface, ResponseInterface) (NetChainElement, error)
)
type NetMiddlewareGroupInterface interface {
	AddPreStreamHandler(fn NetPreStreamHandlerFn)
	AddRequestHandler(fn NetRequestHandlerFn)
	AddResponseHandler(fn NetRequestHandlerFn)
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

type CallInterface interface {
	Request() RequestInterface
	Response() ResponseInterface
	Invoke() (interface{}, error)
}

type NetCallInterface interface {
	CallInterface
	Session() SessionInterface
}

type LocalCallInterface interface {
	CallInterface
}

type SessionInterface interface {
	Id() SessionId
}
