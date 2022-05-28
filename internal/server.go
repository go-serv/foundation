package internal

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
)

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

type LocalServerInterface interface {
	ServerInterface
}

type NetworkServerInterface interface {
	ServerInterface
}
