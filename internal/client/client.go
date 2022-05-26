package client

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/grpc"
	"net"
)

type client struct {
	codec    i.CodecInterface
	endpoint i.EndpointInterface
	conn     net.Conn
	dialOpts []grpc.DialOption
	mwGroup  i.MiddlewareGroupInterface
	insecure bool
	ancillary.MethodMustBeImplemented
}

type localClient struct {
	client
}

type netClient struct {
	client
	svc i.NetworkServiceInterface
}

func (c *client) Codec() i.CodecInterface {
	return c.codec
}

func (c *client) WithCodec(cc i.CodecInterface) {
	c.codec = cc
}

func (c *client) Endpoint() i.EndpointInterface {
	return c.endpoint
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.MethodMustBeImplemented.Panic()
}

func (c *netClient) NetService() i.NetworkServiceInterface {
	return c.svc
}
