package server

import (
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/service/net"
)

func NewCerberusService(eps []z.EndpointInterface, cfg ConfigInterface) z.NetworkServiceInterface {
	svc := new(cerberusService)
	svc.NetworkServiceInterface = net.NewNetworkService(Name, cfg, eps)
	return svc
}
