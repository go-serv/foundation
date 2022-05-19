package client

import "github.com/go-serv/service/internal/server"

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
	return c
}
