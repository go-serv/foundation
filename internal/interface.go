package internal

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ServiceDescriptorInterface interface {
	Descriptor() protoreflect.ServiceDescriptor
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
	AddServiceProtoExt(ext *protoimpl.ExtensionInfo)
	AddMethodProtoExt(ext *protoimpl.ExtensionInfo)
	AddMessageExtension(ext *protoimpl.ExtensionInfo)
	Populate()
	MethodDescriptorByName(protoreflect.FullName) (MethodDescriptorInterface, bool)
}

type MethodDescriptorInterface interface {
	Interface() protoreflect.MethodDescriptor
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
}

type EndpointInterface interface {
	Address() string
	Listen() error
	ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	GrpcServer() *grpc.Server
	WithServer(ServerInterface)
}

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

type SymCipherInterface interface {
	WithNonce([]byte)
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
