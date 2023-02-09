package service

import (
	job "github.com/AgentCoop/go-work"
	"github.com/mesh-master/foundation/pkg/z"
	"google.golang.org/grpc"
)

type endpoint struct {
	appServer      z.AppServerInterface
	services       []z.ServiceInterface
	grpcServer     *grpc.Server
	grpcServerOpts []grpc.ServerOption
}

func (ep *endpoint) AppServer() z.AppServerInterface {
	return ep.appServer
}

func (ep *endpoint) BindAppServer(app z.AppServerInterface) {
	ep.appServer = app
}

func (ep *endpoint) Services() []z.ServiceInterface {
	return ep.services
}

func (ep *endpoint) AddService(svc z.ServiceInterface) {
	ep.services = append(ep.services, svc)
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
