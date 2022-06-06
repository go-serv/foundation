package net

import (
	"context"
	z "github.com/go-serv/service/internal"
	"google.golang.org/grpc"
)

type netContext struct {
	context.Context
	req     z.RequestInterface
	res     z.ResponseInterface
	handler grpc.UnaryHandler
}

func (ctx *netContext) Request() z.RequestInterface {
	return ctx.req
}

func (ctx *netContext) Response() z.ResponseInterface {
	return ctx.res
}

func (ctx *netContext) Invoke() (res interface{}, err error) {
	res, err = ctx.handler(ctx, ctx.req.Payload())
	return
}

func (ctx *netContext) Session() z.SessionInterface {
	return nil
}
