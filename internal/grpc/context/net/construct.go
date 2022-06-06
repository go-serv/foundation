package net

import (
	"context"
	"github.com/go-serv/service/internal/grpc/response/net"
	"github.com/go-serv/service/pkg/z"
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
