package client

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	"github.com/go-serv/service/internal/grpc/codec"
	"google.golang.org/grpc"
	"net"
)

type client struct {
	endpoint i.EndpointInterface
	conn     net.Conn
	dialOpts []grpc.DialOption
	mwGroup  i.MiddlewareGroupInterface
	msgProc  codec.MessageProcessorInterface
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

func (c *client) MessageProcessor() codec.MessageProcessorInterface {
	return c.msgProc
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
