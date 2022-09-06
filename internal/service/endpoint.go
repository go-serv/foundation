package service

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

type endpoint struct {
	svc            z.ServiceInterface
	grpcServer     *grpc.Server
	grpcServerOpts []grpc.ServerOption
}

func (ep *endpoint) Service() z.ServiceInterface {
	return ep.svc
}

func (ep *endpoint) BindService(svc z.ServiceInterface) {
	ep.svc = svc
}

func (ep *endpoint) Address() string {
	return ""
}

func (ep *endpoint) GrpcServer() *grpc.Server {
	return ep.grpcServer
}

func (ep *endpoint) BindGrpcServer(grpc *grpc.Server) {
	ep.grpcServer = grpc
}

func (ep *endpoint) GrpcServerOptions() []grpc.ServerOption {
	return ep.grpcServerOpts
}

func (ep *endpoint) WithGrpcServerOptions(opts ...grpc.ServerOption) {
	ep.grpcServerOpts = append(ep.grpcServerOpts, opts...)
}

func (ep *endpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	return nil, nil, nil
}
