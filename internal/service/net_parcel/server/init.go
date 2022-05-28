package server

import (
	rt "github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service/net"
)

func init() {
	svc := new(netParcel)
	svc.NetworkServiceInterface = net.NewNetworkService(Name)
	rt.Runtime().RegisterNetworkService(svc)
}
