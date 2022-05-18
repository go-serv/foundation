package calldesc

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func newDescriptor(ctx context.Context) callDesc {
	desc := callDesc{}
	desc.Context = ctx
	return desc
}

func NewServerCallDesc(ctx context.Context) (*callDescServer, error) {
	desc := &callDescServer{newDescriptor(ctx)}
	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil
	}
	//desc.req = &request{meta: meta.NewMeta(header)}
	//desc.res = &response{meta: metadata.MD{}}
	return desc, nil
}
