package net

import (
	"context"
	"github.com/go-serv/foundation/pkg/z"
)

type netContext struct {
	context.Context
	service  z.NetworkServiceInterface
	req      z.RequestInterface
	res      z.ResponseInterface
	tenantId z.TenantId
}

func (ctx *netContext) Interface() context.Context {
	return ctx.Context
}

func (ctx *netContext) WithInterface(value context.Context) {
	ctx.Context = value
}

func (ctx *netContext) NetworkService() z.NetworkServiceInterface {
	return ctx.service
}

func (ctx *netContext) Tenant() z.TenantId {
	return ctx.tenantId
}

func (ctx *netContext) WithTenant(id z.TenantId) {
	ctx.tenantId = id
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
