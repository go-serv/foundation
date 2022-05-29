package internal

import (
	"google.golang.org/grpc"
)

type MiddlewareGroupInterface interface {
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	UnaryClientInterceptor() grpc.UnaryClientInterceptor
}

type MetaInterface interface {
}

type RequestInterface interface {
	MethodName() string
	Data() interface{}
	WithData(data interface{})
}

type ResponseInterface interface {
}
