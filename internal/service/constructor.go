package service

import "github.com/mesh-master/foundation/pkg/z"

func NewService(name string, cfg z.ServiceCfgInterface, endpoints []z.EndpointInterface) *service {
	s := new(service)
	s.name = name
	s.cfg = cfg
	s.endpoints = endpoints
	s.state = StateInit
	//app.AddService(s)
	return s
}

func NewEndpoint() *endpoint {
	ep := &endpoint{}
	return ep
}
