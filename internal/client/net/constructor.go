package net

import (
	"github.com/go-serv/foundation/internal/client"
	net_cc "github.com/go-serv/foundation/internal/grpc/codec/net"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	net_mw "github.com/go-serv/foundation/internal/grpc/middleware/net"
	enc_mw "github.com/go-serv/foundation/internal/middleware/net/enc_msg"
	session_mw "github.com/go-serv/foundation/internal/middleware/net/session"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func NewClient(svcName string, e z.EndpointInterface) *netClient {
	c := new(netClient)
	c.insecure = true
	c.ClientInterface = client.NewClient(svcName, e)
	//c.svc = net_service.NewNetworkService(svcName, nil, []z.EndpointInterface{e}, nil)
	c.WithMeta(meta_net.NewMeta(nil))
	// Set client codec
	codec := net_cc.NewOrRegistered(string(svcName))
	c.WithCodec(codec)
	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	//
	c.newDefaultMiddleware()
	runtime.Runtime().RegisterClient(c)
	return c
}

func (c *netClient) newDefaultMiddleware() {
	mw := net_mw.NewMiddleware()
	mw.WithClient(c)
	c.WithMiddleware(mw)
	session_mw.ClientInit(mw)
	enc_mw.ClientInit(mw)
}
