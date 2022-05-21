package net

import (
	"context"
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/grpc"
)

func FromServerContext(ctx context.Context, data interface{}) *netRequest {
	r := &netRequest{}
	r.Context = ctx
	name, _ := grpc.Method(r.Context)
	r.Method = ancillary.GrpcDotNotation(name).MethodName()
	//r.Meta = md
	r.WithData(data)
	return r
}

func FromClientContext(ctx context.Context, data interface{}, methodName string) *netRequest {
	if ctx == nil {
		ctx = context.Background()
	}
	r := &netRequest{}
	r.Context = ctx
	r.Method = ancillary.GrpcDotNotation(methodName).MethodName()
	r.WithData(data)
	return r
}
