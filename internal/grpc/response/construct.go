package response

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func newResponse(ctx context.Context, data interface{}, md metadata.MD) response {
	r := response{}
	r.Context = ctx
	r.data = data
	r.meta = md
	return r
}

func NewNetResponse(ctx context.Context, data interface{}, md metadata.MD) *netResponse {
	r := &netResponse{newResponse(ctx, data, md)}
	return r
}
