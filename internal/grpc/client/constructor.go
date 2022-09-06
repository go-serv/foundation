package client

import (
	"github.com/go-serv/foundation/pkg/z"
)

func NewClient(svcName string, e z.EndpointInterface) *client {
	c := new(client)
	c.svcName = svcName
	c.endpoint = e
	return c
}
