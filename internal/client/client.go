package client

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/grpc"
	"net"
)

type client struct {
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

func (c *client) Client_Endpoint() i.EndpointInterface {
	return c.endpoint
}

func (c *client) Client_NewClient(cc grpc.ClientConnInterface) {
	c.MethodMustBeImplemented.Panic()
}

func (c *netClient) Client_NetService() i.NetworkServiceInterface {
	return c.svc
}
