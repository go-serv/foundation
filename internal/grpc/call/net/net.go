package net

import (
	"context"
	z "github.com/go-serv/service/internal"
	"google.golang.org/grpc"
)

type NetCall struct {
	ctx     context.Context
	req     z.RequestInterface
	res     z.ResponseInterface
	handler grpc.UnaryHandler
}

func (call *NetCall) Request() z.RequestInterface {
	return call.req
}

func (call *NetCall) Response() z.ResponseInterface {
	return call.res
}

func (call *NetCall) Invoke() (res interface{}, err error) {
	res, err = call.handler(call.ctx, call.req.Payload())
	return
}
