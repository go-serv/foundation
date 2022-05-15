package service

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	net_svc "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/logger"
	_ "github.com/go-serv/service/internal/logger"
	"google.golang.org/grpc"
)

type State int

const (
	StateInit State = iota
	StateRunning
	StateStopInProgress
	StateStopped
	StateSuspended // service is running but incoming requests are not being processed
)

type baseService struct {
	// Service name in dot notation
	name           string
	state          State
	cfg            ConfigInterface
	endpoints      []EndpointInterface
	grpcServersJob job.JobInterface
	mustBeImplemented
}

type LocalService struct {
	baseService
}

type networkService struct {
	baseService
	net_svc.NetParcelServer
}

func (s *baseService) Service_AddEndpoint(e EndpointInterface) {
	s.endpoints = append(s.endpoints, e)
}

func (s *baseService) Service_Name(short bool) string {
	return s.name
}

func (s *baseService) Service_State() State {
	return s.state
}

func (s *baseService) Service_Start() {
	if len(s.endpoints) == 0 {
		panic("no service endpoints have been specified")
	}
	// Start listening on baseService endpoints
	for _, e := range s.endpoints {
		s.grpcServersJob.AddTask(e.ServeTask)
	}
	<-s.grpcServersJob.Run()
}

func (s *baseService) Service_Stop() {
	s.state = StateStopInProgress
	info := job.Logger(logger.Info)
	for _, e := range s.endpoints {
		info("stopping %s on %s", s.Service_Name(false), e.Address())
		e.GrpcServer().GracefulStop()
	}
	s.state = StateStopped
}

func (s *networkService) NetParcel() net_svc.NetParcelServer {
	return s.NetParcelServer
}

//
// Methods to implement
//
type mustBeImplemented struct {
}

func (mustBeImplemented) panic(methodName string) {
	panic(fmt.Sprintf("method service.%s must be implemented", methodName))
}

func (m mustBeImplemented) Service_Register(srv *grpc.Server) {
	m.panic("Service_Register")
}

//
