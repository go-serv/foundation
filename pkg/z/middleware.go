package z

import (
	"google.golang.org/grpc"
)

type MiddlewareInterface interface {
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	UnaryClientInterceptor() grpc.UnaryClientInterceptor
	Client() ClientInterface
	WithClient(ClientInterface)
}

type MiddlewareAwareInterface interface {
	Middleware() MiddlewareInterface
	WithMiddleware(mg MiddlewareInterface)
}

type (
	NetPreStreamHandlerFn func(ServiceReflectInterface, MethodReflectionInterface, MessageReflectionInterface) error
	NetRequestHandlerFn   func(next NetChainElementFn, req RequestInterface, res ResponseInterface) error
	NetResponseHandlerFn  func(next NetChainElementFn, res ResponseInterface) error
	NetChainElementFn     func(RequestInterface, ResponseInterface) (NetChainElementFn, error)
)
type NetMiddlewareInterface interface {
	MiddlewareInterface
	AddPreStreamHandler(fn NetPreStreamHandlerFn)
	AddRequestHandler(fn NetRequestHandlerFn)
	AddResponseHandler(fn NetResponseHandlerFn)
}
