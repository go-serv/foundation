package service

import job "github.com/AgentCoop/go-work"

func newBaseService(name string, cfg ConfigInterface) baseService {
	s := baseService{name: name, cfg: cfg}
	s.state = StateInit
	s.grpcServersJob = job.NewJob(nil)
	return s
}

func NewLocalService(name string, cfg ConfigInterface) LocalServiceInterface {
	s := &LocalService{newBaseService(name, cfg)}
	return s
}

func NewNetworkService(name string, cfg ConfigInterface) NetworkServiceInterface {
	s := &NetworkService{newBaseService(name, cfg)}
	return s
}

//
// Endpoints
//

func NewEndpoint(s BaseServiceInterface) endpoint {
	e := endpoint{service: s}
	return e
}

func NewTcpEndpoint(s BaseServiceInterface, hostname string, port int) tcpEndpoint {
	return tcpEndpoint{
		endpoint: NewEndpoint(s),
		hostname: hostname,
		port:     port,
	}
}

func NewTcp4Endpoint(s NetworkServiceInterface, hostname string, port int) *tcp4Endpoint {
	e := &tcp4Endpoint{NewTcpEndpoint(s, hostname, port)}
	return e
}

func NewTcp6Endpoint(s NetworkServiceInterface, hostname string, port int) *tcp6Endpoint {
	e := &tcp6Endpoint{NewTcpEndpoint(s, hostname, port), nil}
	return e
}
