package net

import (
	"context"
	z "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/grpc/response/net"
	"google.golang.org/grpc"
)

func NewCall(ctx context.Context, req z.RequestInterface, handler grpc.UnaryHandler) *netContext {
	c := new(netContext)
	c.Context = ctx
	c.req = req
	c.res = net.NewResponse()
	c.handler = handler
	return c
}
