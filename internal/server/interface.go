package server

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
)

type EndpointInterface interface {
	Address() string
	Listen() error
	ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	GrpcServer() *grpc.Server
	withServer(serverInterface)
}

type serverInterface interface {
	AddGrpcServerOption(opt grpc.ServerOption)
	GrpcServerOptions() []grpc.ServerOption
	AddEndpoint(endpoint EndpointInterface)
	Endpoints() []EndpointInterface
	Start()
	Stop()
	MainJob() job.JobInterface
}

type NetworkServerInterface interface {
	serverInterface
}
