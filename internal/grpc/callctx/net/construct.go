package net

import (
	"context"
	"github.com/go-serv/foundation/internal/grpc/msg/response"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

func NewServerContext(ctx context.Context, req z.RequestInterface, handler grpc.UnaryHandler) *srvContext {
	srvctx := new(srvContext)
	srvctx.Context = ctx
	srvctx.req = req
	res, _ := response.NewResponse(nil, nil)
	srvctx.res = res
	srvctx.handler = handler
	return srvctx
}

func NewClientContext(ctx context.Context) *clientCtx {
	c := new(clientCtx)
	c.Context = ctx
	return c
}
