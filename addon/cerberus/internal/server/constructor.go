package server

import (
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
)

func NewCerberusService(eps []z.EndpointInterface, cfg ConfigInterface) z.NetworkServiceInterface {
	svc := new(cerberusService)
	svc.NetworkServiceInterface = net.NewNetworkService(Name, cfg, eps)
	return svc
}
