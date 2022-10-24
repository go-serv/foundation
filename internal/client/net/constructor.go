package net

import (
	"github.com/go-serv/foundation/internal/client"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	mw_net "github.com/go-serv/foundation/internal/middleware/net"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
)

func NewClient(svcName string, ep z.EndpointInterface) *netClient {
	netCt := new(netClient)
	netCt.ClientInterface = client.NewClient(svcName, ep)
	netCt.svc = net.NewNetworkService(svcName, nil, nil)

	meta := meta_net.NewMeta(nil)
	netCt.WithMeta(meta)

	netCt.Middleware().WithClient(netCt)
	netCt.Middleware().Append(z.NetworkMwKey, mw_net.ClientRequestNetHandler, mw_net.ClientResponseNetHandler)

	runtime.Runtime().RegisterClient(netCt)
	service.Reflection().AddService(svcName)
	// todo: handler error
	service.Reflection().Populate()
	return netCt
}
