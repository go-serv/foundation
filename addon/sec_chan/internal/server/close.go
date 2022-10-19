package server

import (
	"context"
	proto "github.com/go-serv/foundation/internal/autogen/sec_chan"
)

func (s impl) Close(ctx context.Context, req *proto.Close_Request) (res *proto.Close_Response, err error) {
	a := 1
	_ = a
	return
}
