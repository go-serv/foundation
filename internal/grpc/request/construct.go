package request

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func newUnaryRequest(ctx context.Context, data interface{}, md metadata.MD) unaryRequest {
	r := unaryRequest{}
	r.Context = ctx
	r.meta = md
	r.data = data
	return r
}

func NewNetUnaryRequest(ctx context.Context, data interface{}, md metadata.MD) *unaryNetRequest {
	r := &unaryNetRequest{newUnaryRequest(ctx, data, md)}
	return r
}
