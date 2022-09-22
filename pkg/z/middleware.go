package z

import (
	"google.golang.org/grpc"
)

type MiddlewareInterface interface {
	AddPreStreamHandler(fn MiddlewarePreStreamHandlerFn)
	AddRequestHandler(fn MiddlewareRequestHandlerFn)
	AddResponseHandler(fn MiddlewareResponseHandlerFn)
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	UnaryClientInterceptor() grpc.UnaryClientInterceptor
	Client() ClientInterface
	WithClient(ClientInterface)
	MergeWithParent(MiddlewareInterface)
}

type MiddlewareAwareInterface interface {
	Middleware() MiddlewareInterface
	WithMiddleware(mg MiddlewareInterface)
}

type (
	MiddlewarePreStreamHandlerFn func(ServiceReflectInterface, MethodReflectionInterface, MessageReflectionInterface) error
	MiddlewareRequestHandlerFn   func(MiddlewareChainElementFn, NetContextInterface, RequestInterface) error
	MiddlewareResponseHandlerFn  func(MiddlewareChainElementFn, NetContextInterface, ResponseInterface) error
	MiddlewareChainElementFn     func(RequestInterface, ResponseInterface) (MiddlewareChainElementFn, error)
)

type NetMiddlewareInterface interface {
	MiddlewareInterface
	AddPreStreamHandler(fn MiddlewarePreStreamHandlerFn)
	AddRequestHandler(fn MiddlewareRequestHandlerFn)
	AddResponseHandler(fn MiddlewareResponseHandlerFn)
}
