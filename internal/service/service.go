package service

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
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

const Foo = 1

func (s State) String() string {
	return [...]string{"StateInit", "Running", "StopInProgress", "Stopped", "Suspended"}[s]
}

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

type NetworkService struct {
	baseService
}

func (s *baseService) Service_AddEndpoint(e EndpointInterface) {
	s.endpoints = append(s.endpoints, e)
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
	for _, e := range s.endpoints {
		e.GrpcServer().GracefulStop()
	}
	s.state = StateStopped
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
