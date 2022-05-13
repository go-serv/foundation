package service

import job "github.com/AgentCoop/go-work"

func newBaseService(name string, cfg ConfigInterface) baseService {
	s := baseService{name: name, cfg: cfg}
	s.mainJob = job.NewJob(nil)
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
