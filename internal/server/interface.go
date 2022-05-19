package server

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
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
	MiddlewareGroup() mw_group.MiddlewareGroupInterface
	WithMiddlewareGroup(mg mw_group.MiddlewareGroupInterface)
}

type NetworkServerInterface interface {
	serverInterface
}
