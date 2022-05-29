package server

import (
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service/net"
)

func init() {
	svc := new(netParcel)
	svc.NetworkServiceInterface = net.NewNetworkService(Name)
	rt := runtime.Runtime()
	rt.RegisterNetworkService(svc)
	rt.Reflection().AddService(Name)
	rt.Reflection().Populate()
}
