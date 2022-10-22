package net

import (
	"github.com/go-serv/foundation/internal/client"
	net_cc "github.com/go-serv/foundation/internal/grpc/codec/net"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	net_mw "github.com/go-serv/foundation/internal/middleware/net"

	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
	"google.golang.org/grpc"
)

func NewClient(svcName string, e z.EndpointInterface) *netClient {
	netCt := new(netClient)
	netCt.ClientInterface = client.NewClient(svcName, e)
	netCt.svc = net.NewNetworkService(svcName, nil, nil)
	meta := meta_net.NewMeta(nil)
	netCt.WithMeta(meta)
	netCt.Middleware().WithClient(netCt)
	netCt.Middleware().Append(z.NetworkMwKey, net_mw.ClientRequestNetHandler, net_mw.ClientResponseNetHandler)
	// Set client codec
	codec := net_cc.NewOrRegistered("proto")
	//netCt.WithCodec(codec)
	netCt.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	//
	//netCt.newDefaultMiddleware()
	runtime.Runtime().RegisterClient(netCt)
	service.Reflection().AddService(svcName)
	service.Reflection().Populate()
	return netCt
}

func (c *netClient) newDefaultMiddleware() {
	//mw := net_mw.NewMiddleware()
	//mw.WithClient(c)
	//c.WithMiddleware(mw)
	//session_mw.ClientInit(mw)
	//sec_chan.MiddlewareClientInit(mw)
}
