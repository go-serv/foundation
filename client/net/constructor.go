package net

import (
	"github.com/go-serv/foundation/internal/grpc/client"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	mw_net "github.com/go-serv/foundation/internal/middleware/net"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
)

func NewClient(svcName string, ep z.EndpointInterface) (nc *netClient, err error) {
	nc = new(netClient)
	nc.ClientInterface = client.NewClient(svcName, ep)
	nc.svc = net.NewNetworkService(svcName, nil, nil)

	meta := meta_net.NewRequestMeta(nil)
	nc.WithMeta(meta)

	nc.Middleware().WithClient(nc)
	nc.Middleware().Append(z.NetworkMwKey, mw_net.ClientRequestNetHandler, mw_net.ClientResponseNetHandler)

	// Set API key.
	var rv any
	if rv, err = runtime.Runtime().Resolve(z.ApiKeyResolver); err != nil {
		return
	}
	if rv != nil {
		nc.WithApiKey(rv.([]byte))
	}

	runtime.Runtime().RegisterClient(nc)
	service.Reflection().AddService(svcName)
	// todo: handler error
	err = service.Reflection().Populate()
	return
}
