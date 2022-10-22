package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/logger"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func (ep *tcpEndpoint) serviceStartInfo(name string) {
	info := job.Logger(logger.Info)
	var extra string
	if ep.IsSecure() {
		extra = "gRPC, with TLS"
	} else if ep.webProxy != nil {
		extra = "gRPC web proxy"
	} else {
		extra = "gRPC, no TLS"
	}
	info("%s started to serve network requests on %s ( %s )", name, ep.Address(), extra)
}

func (ep *tcpEndpoint) initTask() {
	//	svc := ep.Service()

	// An option for the middleware unary interceptor.
	mwUnaryInt := ep.AppServer().Middleware().UnaryServerInterceptor()
	ep.WithGrpcServerOptions(grpc.ChainUnaryInterceptor(mwUnaryInt))

	// Codec option.
	//if svc.Codec() != nil {
	//	ep.WithGrpcServerOptions(grpc.ForceServerCodec(svc.Codec()))
	//}

	grpcServer := grpc.NewServer(ep.GrpcServerOptions()...)
	ep.BindGrpcServer(grpcServer)
	for _, svc := range ep.Services() {
		svc.(z.NetworkServiceInterface).Register(ep.GrpcServer())
		svc.(z.NetworkServiceInterface).WithTlsEnabled(ep.IsSecure())
		ep.serviceStartInfo(svc.Name())
	}
}

func (ep *tcp4Endpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		ep.initTask()
	}
	run := func(task job.TaskInterface) {
		err := ep.listenAndServe()
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}

func (ep *tcp6Endpoint) ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize) {
	init := func(task job.TaskInterface) {
		ep.initTask()
	}
	run := func(task job.TaskInterface) {
		err := ep.listenAndServe()
		task.Assert(err)
		task.Done()
	}
	fin := func(task job.TaskInterface) {
	}
	return init, run, fin
}
