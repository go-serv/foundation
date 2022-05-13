package service

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
)

type baseService struct {
	// Service name in dot notation
	name      string
	cfg       ConfigInterface
	endpoints []EndpointInterface
	mainJob   job.JobInterface
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

func (s *baseService) Service_Start() {
	if len(s.endpoints) == 0 {
		panic("no service endpoints have been specified")
	}
	// Start listening on baseService endpoints
	for _, e := range s.endpoints {
		s.mainJob.AddTask(e.ServeTask)
	}
	<-s.mainJob.Run()
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
