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
	WithServer(serverInterface)
}

type serverInterface interface {
	AddEndpoint(endpoint EndpointInterface)
	Start()
	Stop()
}

type NetworkServerInterface interface {
	serverInterface
}
