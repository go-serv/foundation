package internal

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/grpc/descriptor"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

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

type ServerInterface interface {
	AddGrpcServerOption(opt grpc.ServerOption)
	GrpcServerOptions() []grpc.ServerOption
	AddEndpoint(endpoint EndpointInterface)
	Endpoints() []EndpointInterface
	Start()
	Stop()
	MainJob() job.JobInterface
	MiddlewareGroup() MiddlewareGroupInterface
	WithMiddlewareGroup(mg MiddlewareGroupInterface)
}

type NetworkServerInterface interface {
	ServerInterface
}

type clientInterface interface {
	Client_Endpoint() EndpointInterface
	Client_NewClient(cc grpc.ClientConnInterface)
	Client_ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
}

type NetworkClientInterface interface {
	Client_Endpoint() EndpointInterface
	Client_NewClient(cc grpc.ClientConnInterface)
	Client_ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	Client_NetService() NetworkServiceInterface
}

type BaseServiceInterface interface {
	Service_Descriptor() descriptor.ServiceDescriptorInterface
	Service_AddServiceProtoExtension(info *protoimpl.ExtensionInfo)
	Service_AddMethodProtoExtension(info *protoimpl.ExtensionInfo)
	Service_Register(srv *grpc.Server)
}

type NetworkServiceInterface interface {
	BaseServiceInterface
	Service_OnNewSession(req RequestInterface) error
	// Service_InfoNewSession returns timeout in seconds for a new session. Zero means no new session is required
	Service_InfoNewSession(methodName string) int32
	Service_InfoMsgEncryption(methodName string) bool

	Service_EncriptionKey() []byte
	Service_WithEncriptionKey([]byte)
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
