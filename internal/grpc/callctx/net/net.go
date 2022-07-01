package net

import (
	"context"
	"github.com/go-serv/service/pkg/z"
)

type netContext struct {
	context.Context
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
