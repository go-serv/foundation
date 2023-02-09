package client

import (
	"github.com/mesh-master/foundation/internal/grpc/middleware"
	"github.com/mesh-master/foundation/internal/middleware/session"
	"github.com/mesh-master/foundation/pkg/z"
)

func NewClient(svcName string, e z.EndpointInterface) *client {
	c := new(client)
	c.svcName = svcName
	c.endpoint = e
	c.mw = middleware.NewClientMiddleware()
	c.mw.Append(z.SessionMwKey, session.ClientRequestSessionHandler, nil)
	return c
}
