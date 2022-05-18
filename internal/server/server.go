package server

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/internal/logger"
)

const (
	Tcp4Network = "tcp4"
	Tcp6Network = "tcp6"
	dsNetwork   = "unixpacket"
)

type Server struct {
	endpoints []EndpointInterface
	grpcJob   job.JobInterface
}

type localServer struct {
	Server
}

func (s *Server) AddEndpoint(e EndpointInterface) {
	e.WithServer(s)
	s.endpoints = append(s.endpoints, e)
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
