package internal

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type MiddlewareGroupInterface interface {
	UnaryServerInterceptor() grpc.UnaryServerInterceptor
	UnaryClientInterceptor() grpc.UnaryClientInterceptor
}

type MetaInterface interface {
	Hydrate(MetaDictionary) error
	Dehydrate(dictionary MetaDictionary) error
}

// Implemented by a Go struct with the public fields and information about headers in field tags
type MetaDictionary interface{}

type RequestInterface interface {
	proto.Message
	Data() interface{}
	WithData(data interface{})
}

type ResponseInterface interface {
	proto.Message
	ToGrpc() interface{}
}
