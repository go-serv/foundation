package net

import (
	"context"
	z "github.com/go-serv/service/internal"
	"google.golang.org/grpc"
)

type NetContext struct {
	context.Context
	req     z.RequestInterface
	res     z.ResponseInterface
	handler grpc.UnaryHandler
}

func (ctx *NetContext) Request() z.RequestInterface {
	return ctx.req
}

func (ctx *NetContext) Response() z.ResponseInterface {
	return ctx.res
}

func (ctx *NetContext) Invoke() (res interface{}, err error) {
	res, err = ctx.handler(ctx, ctx.req.Payload())
	return
}

func (ctx *NetContext) Session() z.SessionInterface {
	return nil
}
