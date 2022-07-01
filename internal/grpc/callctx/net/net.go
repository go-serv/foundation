package net

import (
	"github.com/go-serv/service/internal/grpc/callctx"
	"github.com/go-serv/service/internal/grpc/codec"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type netContext struct {
	callctx.CallCtx
	req      z.RequestInterface
	res      z.ResponseInterface
	tenantId z.TenantId
}

func (ctx *netContext) Tenant() z.TenantId {
	return ctx.tenantId
}

func (ctx *netContext) WithTenant(id z.TenantId) {
	ctx.tenantId = id
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

func (ctx *srvContext) Invoke() (err error) {
	var (
		ifreplay any
	)
	if ifreplay, err = ctx.handler(ctx, ctx.req.DataFrame().Interface()); err != nil {
		return
	}
	msg := ifreplay.(proto.Message)
	ctx.res.WithDataFrame(codec.NewDataFrame(msg))
	err = ctx.res.Populate(msg)
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
func (ctx *clientContext) Invoke() (err error) {
	var (
		md metadata.MD
	)
	if md, err = ctx.req.Meta().Dehydrate(); err != nil {
		return
	}
	ctx.Context = metadata.NewOutgoingContext(ctx.Context, md)
	methodReflect := ctx.req.MethodReflection()
	methodName := methodReflect.SlashFullName()
	if err = ctx.invoker(
		ctx,
		methodName,
		ctx.req.DataFrame(),
		ctx.res.DataFrame(),
		ctx.cc,
		ctx.opts...); err != nil {
		return
	}
	// Response meta data.
	if err = ctx.res.Meta().Hydrate(); err != nil {
		return
	}
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
