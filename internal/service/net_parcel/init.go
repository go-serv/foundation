package net_parcel

import (
	rt "github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service"
)

func init() {
	svc := new(netParcel)
	svc.NetworkServiceInterface = service.NewNetworkService(Name)
	rt.Runtime().RegisterNetworkService(Name, svc)
}
