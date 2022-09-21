package net

import (
	"github.com/go-serv/foundation/addon/sec-chan-mw"
	"github.com/go-serv/foundation/internal/client"
	net_cc "github.com/go-serv/foundation/internal/grpc/codec/net"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	net_mw "github.com/go-serv/foundation/internal/grpc/middleware/net"
	//enc_mw "github.com/go-serv/foundation/internal/middleware/net/enc_msg"
	session_mw "github.com/go-serv/foundation/internal/middleware/net/session"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
	"google.golang.org/grpc"
)

func NewClient(svcName string, e z.EndpointInterface) *netClient {
	c := new(netClient)
	c.ClientInterface = client.NewClient(svcName, e)
	c.svc = net.NewNetworkService(svcName, nil, nil)
	meta := meta_net.NewMeta(nil)
	c.WithMeta(meta)
	// Set client codec
	codec := net_cc.NewOrRegistered("proto")
	c.WithCodec(codec)
	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	//
	c.newDefaultMiddleware()
	runtime.Runtime().RegisterClient(c)
	service.Reflection().AddService(svcName)
	service.Reflection().Populate()
	return c
}

func (c *netClient) newDefaultMiddleware() {
	mw := net_mw.NewMiddleware()
	mw.WithClient(c)
	c.WithMiddleware(mw)
	session_mw.ClientInit(mw)
	sec_chan_mw.MiddlewareClientInit(mw)
}
