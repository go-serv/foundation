package internal

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ServiceDescriptorInterface interface {
	Descriptor() protoreflect.ServiceDescriptor
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
	AddServiceProtoExt(ext *protoimpl.ExtensionInfo)
	AddMethodProtoExt(ext *protoimpl.ExtensionInfo)
	Populate()
	MethodDescriptorByName(protoreflect.FullName) (MethodDescriptorInterface, bool)
}

type MethodDescriptorInterface interface {
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

type HeaderFlags32Type uint32

func (f HeaderFlags32Type) Has(chkFlag HeaderFlags32Type) bool {
	return f&chkFlag != 0
}

type DataFrameInterface interface {
	Parse([]byte) error
	ParseHook() error
	HeaderFlags() HeaderFlags32Type
	WithHeaderFlag(HeaderFlags32Type)
	Compose() ([]byte, error)
	Payload() []byte
	WithPayload([]byte)
	AttachData(b []byte)
}

type MsgProcTaskHandler func(next MsgProcTaskHandler, in []byte, md MethodDescriptorInterface, df DataFrameInterface) ([]byte, error)

type MessageProcessTaskInterface interface {
	Execute() ([]byte, error)
}

type MessageProcessorInterface interface {
	AddHandlers(pre MsgProcTaskHandler, post MsgProcTaskHandler)
	NewPreTask(wire []byte, msg proto.Message) (MessageProcessTaskInterface, error)
	NewPostTask(wire []byte, msg proto.Message) (MessageProcessTaskInterface, error)
}

type CodecInterface interface {
	encoding.Codec
	Processor() MessageProcessorInterface
	NewDataFrame() DataFrameInterface
}

type NetworkServerInterface interface {
	ServerInterface
}

type clientInterface interface {
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
	Endpoint() EndpointInterface
	ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	NewClient(cc grpc.ClientConnInterface)
}

type NetworkClientInterface interface {
	clientInterface
	NetService() NetworkServiceInterface
}

type BaseServiceInterface interface {
	Service_Descriptor() ServiceDescriptorInterface
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

type LocalServiceInterface interface {
	BaseServiceInterface
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
