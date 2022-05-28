package net

import i "github.com/go-serv/service/internal"

type netClient struct {
	i.ClientInterface
	svc      i.NetworkServiceInterface
	insecure bool
}

func (c *netClient) NetService() i.NetworkServiceInterface {
	return c.svc
}
