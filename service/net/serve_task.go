package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/logger"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func (ep *tcpEndpoint) initTask() {
	svc := ep.Service()

	// An option for the middleware unary interceptor.
	mwUnaryInt := svc.Middleware().UnaryServerInterceptor()
	ep.WithGrpcServerOptions(grpc.ChainUnaryInterceptor(mwUnaryInt))

	// Codec option.
	if svc.Codec() != nil {
		ep.WithGrpcServerOptions(grpc.ForceServerCodec(svc.Codec()))
	}
	grpcServer := grpc.NewServer(ep.GrpcServerOptions()...)

	ep.BindGrpcServer(grpcServer)
	ep.Service().(z.NetworkServiceInterface).Register(ep.GrpcServer())

	info := job.Logger(logger.Info)
	var extra string
	if ep.IsSecure() {
		extra = "gRPC, with TLS"
	} else if ep.webProxy != nil {
		extra = "gRPC web proxy"
	} else {
		extra = "gRPC, no TLS"
	}
	info("serving network requests on %s ( %s )", ep.Address(), extra)
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
