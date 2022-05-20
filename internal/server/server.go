package server

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/logger"
	"google.golang.org/grpc"
)

const (
	Tcp4Network = "tcp4"
	Tcp6Network = "tcp6"
	dsNetwork   = "unixpacket"
)

type Server struct {
	endpoints []i.EndpointInterface
	mainJob   job.JobInterface
	mwGroup   i.MiddlewareGroupInterface
	grpcOpts  []grpc.ServerOption
}

type localServer struct {
	Server
}

func (s *Server) AddEndpoint(e i.EndpointInterface) {
	e.WithServer(s)
	s.endpoints = append(s.endpoints, e)
}

func (s *Server) Endpoints() []i.EndpointInterface {
	out := make([]i.EndpointInterface, 1)
	for _, e := range s.endpoints {
		out = append(out, e)
	}
	return out
}

func (s *Server) AddGrpcServerOption(opt grpc.ServerOption) {
	s.grpcOpts = append(s.grpcOpts, opt)
}

func (s *Server) GrpcServerOptions() []grpc.ServerOption {
	return s.grpcOpts
}

func (s *Server) Start() {
	if len(s.endpoints) == 0 {
		panic("no Server endpoints have been specified")
	}
	// Start listening on baseService endpoints
	for _, e := range s.endpoints {
		s.mainJob.AddTask(e.ServeTask)
	}
	<-s.mainJob.Run()
}

func (s *Server) MainJob() job.JobInterface {
	return s.mainJob
}

func (s *Server) MiddlewareGroup() i.MiddlewareGroupInterface {
	return s.mwGroup
}

func (s *Server) WithMiddlewareGroup(mw i.MiddlewareGroupInterface) {
	s.mwGroup = mw
}

func (s *Server) Stop() {
	//s.state = StateStopInProgress
	info := job.Logger(logger.Info)
	for _, e := range s.endpoints {
		info("stopping network Server on %s", e.Address())
		e.GrpcServer().GracefulStop()
	}
	//s.state = StateStopped
}
