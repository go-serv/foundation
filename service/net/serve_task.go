package net

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/internal/logger"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func (ep *tcpEndpoint) initTask() {
	// Add a server option for the middleware unary interceptor.
	mwUnaryInt := ep.Service().MiddlewareGroup().UnaryServerInterceptor()
	ep.WithGrpcServerOptions(grpc.ChainUnaryInterceptor(mwUnaryInt))
	grpcServer := grpc.NewServer(ep.GrpcServerOptions()...)
	ep.BindGrpcServer(grpcServer)
	ep.Service().(z.NetworkServiceInterface).Register(ep.GrpcServer())
	info := job.Logger(logger.Info)
	info("serving network requests on %s", ep.Address())
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
