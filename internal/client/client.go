package client

import (
	"github.com/go-serv/service/internal/ancillary"
	"github.com/go-serv/service/internal/server"
	"google.golang.org/grpc"
	"net"
)

type client struct {
	endpoint server.EndpointInterface
	conn     net.Conn
	dialOpts []grpc.DialOption
	insecure bool
	ancillary.MethodMustBeImplemented
}

type localClient struct {
	client
}

type netClient struct {
	client
}

func (c *client) Client_Endpoint() server.EndpointInterface {
	return c.endpoint
}

func (c *client) Client_NewClient(cc grpc.ClientConnInterface) {
	c.MethodMustBeImplemented.Panic()
}
