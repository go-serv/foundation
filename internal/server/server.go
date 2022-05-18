package server

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/logger"
	"google.golang.org/grpc"
)

const (
	Tcp4Network = "tcp4"
	Tcp6Network = "tcp6"
	dsNetwork   = "unixpacket"
)

type Server struct {
	endpoints []EndpointInterface
	grpcJob   job.JobInterface
	grpcOpts  []grpc.ServerOption
}

type localServer struct {
	Server
}

func (s *Server) AddEndpoint(e EndpointInterface) {
	e.withServer(s)
	s.endpoints = append(s.endpoints, e)
}

func (s *Server) Endpoints() []EndpointInterface {
	out := make([]EndpointInterface, 1)
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
		s.grpcJob.AddTask(e.ServeTask)
	}
	<-s.grpcJob.Run()
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
