package client

import (
	"github.com/go-serv/service/internal/grpc/codec/net"
	"github.com/go-serv/service/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

func newClient(e server.EndpointInterface) client {
	c := client{}
	c.endpoint = e
	c.insecure = true
	return c
}

func NewLocalClient(e server.EndpointInterface) *localClient {
	c := &localClient{newClient(e)}
	return c
}

func NewNetClient(e server.EndpointInterface) *netClient {
	c := &netClient{newClient(e)}
	netCodec := encoding.GetCodec(codec.Name)
	c.dialOpts = append(c.dialOpts, grpc.WithDefaultCallOptions(grpc.ForceCodec(netCodec)))
	return c
}
