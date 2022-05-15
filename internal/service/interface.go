package service

import (
	job "github.com/AgentCoop/go-work"
	net_svc "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
)

type EndpointInterface interface {
	Address() string
	Listen() error
	ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	GrpcServer() *grpc.Server
}

type ConfigInterface interface {
}

type GrpcMessageWrapperFn func(in []byte) []byte

// All method names are prefixed with Service_ to avoid name conflicts with the method names of a GRPC service.
type BaseServiceInterface interface {
	Service_Name(bool) string
	Service_Register(srv *grpc.Server)
	Service_AddEndpoint(endpoint EndpointInterface)
	Service_Start()
	Service_Stop()
	Service_State() State
	// Adds a new wrapper to the wrapper chain
	// AddGrpcMessageWrapper(GrpcMessageWrapperFn)
}

type LocalServiceInterface interface {
	BaseServiceInterface
}

type NetworkServiceInterface interface {
	BaseServiceInterface
	NetParcel() net_svc.NetParcelServer
}
