package service

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
)

type EndpointInterface interface {
	Listen() error
	ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
}

type ConfigInterface interface {
}

type GrpcMessageWrapperFn func(in []byte) []byte

// All method names are prefixed with Service_ to avoid name conflicts with the method names of a GRPC service.
type BaseServiceInterface interface {
	Service_Register(srv *grpc.Server)
	Service_AddEndpoint(endpoint EndpointInterface)
	Service_Start()
	// Adds a new wrapper to the wrapper chain
	// AddGrpcMessageWrapper(GrpcMessageWrapperFn)
}

type LocalServiceInterface interface {
	BaseServiceInterface
}

type NetworkServiceInterface interface {
	BaseServiceInterface
}
