package client

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/grpc/codec"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	md_cipher "github.com/go-serv/service/internal/grpc/middleware/cipher_msg"
	mw_net "github.com/go-serv/service/internal/grpc/middleware/mw_group/net"
	net_service "github.com/go-serv/service/internal/service/net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

func newClient(e i.EndpointInterface) client {
	c := client{}
	c.endpoint = e
	c.insecure = true
	return c
}

func NewLocalClient(e i.EndpointInterface) *localClient {
	c := &localClient{newClient(e)}
	return c
}

func NewNetClient(svcName string, e i.EndpointInterface) *netClient {
	c := &netClient{}
	c.client = newClient(e)
	c.svc = net_service.NewNetworkService(svcName)
	// Create message post- and pre-processor
	netCodec := encoding.GetCodec(net_cc.Name).(codec.CodecInterface)
	c.msgProc = codec.NewProcessor(netCodec)
	// Create default group of the client middlewares
	c.mwGroup = c.defaultMiddlewareGroup()
	c.dialOpts = append(c.dialOpts,
		grpc.WithDefaultCallOptions(grpc.ForceCodec(netCodec)),
		grpc.WithChainUnaryInterceptor(c.mwGroup.UnaryClientInterceptor()),
	)
	return c
}

func (c *netClient) defaultMiddlewareGroup() i.MiddlewareGroupInterface {
	g := mw_net.NewMiddlewareGroup(c)
	g.AddItem(md_cipher.NewNetCipherClientHandler(c))
	return g
}
