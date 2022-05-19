package client

import (
	"github.com/go-serv/service/internal/server"
	"google.golang.org/grpc"
	"net"
)

type client struct {
	endpoint server.EndpointInterface
	conn     net.Conn
	dialOpts []grpc.DialOption
}

type localClient struct {
	client
}

type netClient struct {
	client
}
