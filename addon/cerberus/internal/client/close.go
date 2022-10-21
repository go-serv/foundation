package client

import (
	proto "github.com/go-serv/foundation/internal/autogen/net/sec_chan"
)

func (i impl) Close(in *proto.Close_Request) (res *proto.Close_Response, err error) {
	ctx := i.PrepareContext()
	res, err = i.c.grpcClient.Close(ctx, in)
	return
}
