package net

import (
	"context"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type netContext struct {
	context.Context
	req z.RequestInterface
	res z.ResponseInterface
}

type srvContext struct {
	netContext
	handler grpc.UnaryHandler
}

type clientContext struct {
	netContext
	responseMd *metadata.MD
	reqMeta    z.MetaInterface
	cc         *grpc.ClientConn
	invoker    grpc.UnaryInvoker
	opts       []grpc.CallOption
}

func (ctx *netContext) Request() z.RequestInterface {
	return ctx.req
}

func (ctx *netContext) Response() z.ResponseInterface {
	return ctx.res
}

func (ctx *srvContext) Invoke() (res interface{}, err error) {
	res, err = ctx.handler(ctx, ctx.req.Payload())
	return
}

// Invoke prepares metadata and calls gRPC method
func (ctx *clientContext) Invoke() (res interface{}, err error) {
	var (
		md metadata.MD
	)
	md, err = ctx.req.Meta().Dehydrate()
	if err != nil {
		return
	}
	ctx.Context = metadata.NewOutgoingContext(ctx.Context, md)
	methodReflect := ctx.req.MethodReflection()
	methodName := methodReflect.SlashFullName()
	err = ctx.invoker(ctx, methodName, ctx.req.Payload(), ctx.res.Payload(), ctx.cc, ctx.opts...)
	return
}

func (ctx *netContext) Session() z.SessionInterface {
	return nil
}
