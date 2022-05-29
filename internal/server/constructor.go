package server

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/go-serv/service/internal"
)

func NewEndpoint() endpoint {
	e := endpoint{}
	return e
}

func newTcpEndpoint(hostname string, port int) tcpEndpoint {
	e := tcpEndpoint{
		endpoint: NewEndpoint(),
		hostname: hostname,
		port:     port,
	}
	return e
}

func NewTcp4Endpoint(hostname string, port int) *tcp4Endpoint {
	e := &tcp4Endpoint{newTcpEndpoint(hostname, port)}
	return e
}

func NewTcp6Endpoint(hostname string, port int) *tcp6Endpoint {
	e := &tcp6Endpoint{newTcpEndpoint(hostname, port), nil}
	return e
}

func NewLocalEndpoint(svc i.LocalServiceInterface) *localEndpoint {
	pathname := string(svc.Name())
	e := &localEndpoint{NewEndpoint(), pathname}
	return e
}

func NewServer() *server {
	s := new(server)
	s.mainJob = job.NewJob(nil)
	return s
}
