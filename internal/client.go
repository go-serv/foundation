package internal

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ClientInterface interface {
	ServiceName() protoreflect.FullName
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
	CodecMiddlewareGroup() CodecMiddlewareGroupInterface
	Endpoint() EndpointInterface
	ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	NewClient(cc grpc.ClientConnInterface)
	WithDialOption(grpc.DialOption)
	DialOptions() []grpc.DialOption
}

type NetworkClientInterface interface {
	ClientInterface
	NetService() NetworkServiceInterface
}

type LocalClientInterface interface {
	ClientInterface
}
