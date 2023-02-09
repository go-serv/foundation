package server

import (
	"context"
	proto "github.com/mesh-master/foundation/internal/autogen/net/sec_chan"
)

func (s impl) Close(ctx context.Context, req *proto.Close_Request) (res *proto.Close_Response, err error) {
	a := 1
	_ = a
	return
}
