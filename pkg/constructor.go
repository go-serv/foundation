package pkg

import (
	i "github.com/go-serv/service/internal"
	srv "github.com/go-serv/service/internal/server"
)

//func NewNetworkService(name string, cfg ServiceConfig) NetworkServiceInterface {
//	return .NewNetworkService(name, cfg)
//}

func NewTcp4Endpoint(hostname string, port int) i.EndpointInterface {
	return srv.NewTcp4Endpoint(hostname, port)
}

func NewTcp6Endpoint(hostname string, port int) i.EndpointInterface {
	return srv.NewTcp6Endpoint(hostname, port)
}

//func NewLocalEndpoint(svc LocalServiceInterface, pathname string) server.EndpointInterface {
//	return local.NewLocalEndpoint(svc, pathname)
//}
