package net

import "github.com/go-serv/service/pkg/z"

type netClient struct {
	z.ClientInterface
	svc      z.NetworkServiceInterface
	insecure bool
}

func (c *netClient) NetService() z.NetworkServiceInterface {
	return c.svc
}
