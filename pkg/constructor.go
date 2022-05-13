package pkg

import "github.com/go-serv/service/internal/service"

func NewNetworkService(name string, cfg ServiceConfig) NetworkServiceInterface {
	return service.NewNetworkService(name, cfg)
}

func NewTcp4Endpoint(svc NetworkServiceInterface, hostname string, port int) service.EndpointInterface {
	return service.NewTcp4Endpoint(svc, hostname, port)
}

func NewTcp6Endpoint(svc NetworkServiceInterface, hostname string, port int) service.EndpointInterface {
	return service.NewTcp6Endpoint(svc, hostname, port)
}
