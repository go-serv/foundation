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
	server  z.NetworkServerInterface
	sess    z.SessionInterface
	handler grpc.UnaryHandler
}

type clientContext struct {
	netContext
	client  z.NetworkClientInterface
	cc      *grpc.ClientConn
	invoker grpc.UnaryInvoker
	opts    []grpc.CallOption
}

func (ctx *netContext) Request() z.RequestInterface {
	return ctx.req
}

func (ctx *netContext) WithRequest(req z.RequestInterface) {
	ctx.req = req
}

func (ctx *netContext) Response() z.ResponseInterface {
	return ctx.res
}

func (ctx *netContext) WithResponse(res z.ResponseInterface) {
	ctx.res = res
}

func (ctx *srvContext) Invoke() (res interface{}, err error) {
	res, err = ctx.handler(ctx, ctx.req.Payload())
	return
}

func (ctx *clientContext) WithClientInvoker(invoker grpc.UnaryInvoker, cc *grpc.ClientConn, opts []grpc.CallOption) {
	ctx.invoker = invoker
	ctx.cc = cc
	ctx.opts = opts
}

func (ctx *clientContext) Client() z.NetworkClientInterface {
	return ctx.client
}

func (ctx *clientContext) WithClient(client z.NetworkClientInterface) {
	ctx.client = client
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
	if err != nil {
		return
	}
	// Response meta data
	err = ctx.res.Meta().Hydrate()
	if err != nil {
		return
	}
	// Response payload
	res = ctx.res.Payload()
	return
}

func (s *srvContext) Session() z.SessionInterface {
	return s.sess
}

func (s *srvContext) WithSession(sess z.SessionInterface) {
	s.sess = sess
}

func (s *srvContext) Server() z.NetworkServerInterface {
	return s.server
}

func (s *srvContext) WithServer(srv z.NetworkServerInterface) {
	s.server = srv
}
