package server

import (
	job "github.com/AgentCoop/go-work"
	mw_codec "github.com/go-serv/service/internal/grpc/middleware/codec"
	"github.com/go-serv/service/internal/logger"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
)

const (
	Tcp4Network      = "tcp4"
	Tcp6Network      = "tcp6"
	UnixDomainSocket = "unix"
)

type server struct {
	codec        z.CodecInterface
	codecMwGroup z.CodecMiddlewareGroupInterface
	endpoints    []z.EndpointInterface
	mainJob      job.JobInterface
	mwGroup      z.MiddlewareInterface
	grpcOpts     []grpc.ServerOption
}

func (s *server) AddEndpoint(e z.EndpointInterface) {
	e.WithServer(s)
	s.endpoints = append(s.endpoints, e)
}

func (s *server) Endpoints() []z.EndpointInterface {
	out := make([]z.EndpointInterface, 1)
	for _, e := range s.endpoints {
		out = append(out, e)
	}
	return out
}

func (s *server) AddGrpcServerOption(opt grpc.ServerOption) {
	s.grpcOpts = append(s.grpcOpts, opt)
}

func (s *server) GrpcServerOptions() []grpc.ServerOption {
	return s.grpcOpts
}

func (s *server) Start() {
	if len(s.endpoints) == 0 {
		panic("no server endpoints have been specified")
	}
	// Start listening on baseService endpoints
	for _, e := range s.endpoints {
		s.mainJob.AddTask(e.ServeTask)
	}
	<-s.mainJob.Run()
}

func (s *server) MainJob() job.JobInterface {
	return s.mainJob
}

func (s *server) MiddlewareGroup() z.MiddlewareInterface {
	return s.mwGroup
}

func (s *server) WithMiddlewareGroup(mw z.MiddlewareInterface) {
	s.mwGroup = mw
}

func (s *server) Stop() {
	//s.state = StateStopInProgress
	info := job.Logger(logger.Info)
	for _, e := range s.endpoints {
		info("stopping network server on %s", e.Address())
		e.GrpcServer().GracefulStop()
	}
	//s.state = StateStopped
}

func (c *server) Codec() z.CodecInterface {
	return c.codec
}

func (s *server) WithCodec(cc z.CodecInterface) {
	s.codec = cc
	s.codecMwGroup = mw_codec.NewCodecMiddlewareGroup(cc)
}

func (s *server) CodecMiddlewareGroup() z.CodecMiddlewareGroupInterface {
	return s.codecMwGroup
}
