package client

import (
	"github.com/go-serv/foundation/internal/grpc/middleware"
	"github.com/go-serv/foundation/pkg/z"
)

func NewClient(svcName string, e z.EndpointInterface) *client {
	c := new(client)
	c.svcName = svcName
	c.endpoint = e
	c.mw = middleware.NewClientMiddleware()
	return c
}
