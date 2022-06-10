package net

import (
	"context"
	res_net "github.com/go-serv/service/internal/grpc/response/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
)

func NewServerContext(ctx context.Context, req z.RequestInterface, handler grpc.UnaryHandler) *srvContext {
	c := new(srvContext)
	c.Context = ctx
	c.req = req
	c.res = res_net.NewResponse(nil, nil)
	c.handler = handler
	return c
}

func NewClientContext(ctx context.Context) *clientContext {
	c := new(clientContext)
	c.Context = ctx
	return c
}
