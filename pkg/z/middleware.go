package z

import (
	"google.golang.org/grpc"
)

type sessionMiddlewareKey struct{}

var (
	SessionMwKey = sessionMiddlewareKey{}
)

type InsertOp int

const (
	InsertBefore InsertOp = iota
	InsertAfter
)

type (
	RequestHandlerFn  func(NextHandlerFn, NetContextInterface, RequestInterface) error
	ResponseHandlerFn func(NextHandlerFn, NetContextInterface, ResponseInterface) error
	NextHandlerFn     func(RequestInterface, ResponseInterface) (NextHandlerFn, error)
)

type MiddlewareInterface interface {
	//
	Insert(targetKey any, op InsertOp, newKey any, req RequestHandlerFn, res ResponseHandlerFn)
	// Append adds a new element to the end of the application-level chain.
	Append(newKey any, req RequestHandlerFn, res ResponseHandlerFn)
	// AppendToServiceChain adds an element to a service middleware chain.
	// Data flow: req -> application-level chain -> service chain (request handler) -> gRPC call
	// 	-> service chain (response handler) -> application-level chain
	AppendToServiceChain(service string, newKey any, req RequestHandlerFn, res ResponseHandlerFn)
}

type ServerMiddlewareInterface interface {
	MiddlewareInterface
	//
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
}

type ClientMiddlewareInterface interface {
	MiddlewareInterface
	// UnaryClientInterceptor returns the client unary gRPC call interceptor.
	UnaryClientInterceptor() grpc.UnaryClientInterceptor
	// Client returns a bound client instance.
	Client() ClientInterface
	// WithClient binds a client instance to the middleware chain.
	WithClient(ClientInterface)
}

type MiddlewareAwareInterface interface {
	Middleware() MiddlewareInterface
	WithMiddleware(mg MiddlewareInterface)
}
