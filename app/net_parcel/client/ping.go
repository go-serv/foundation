package client

import (
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	net_cc "github.com/mesh-master/foundation/internal/grpc/client"
)

type PingOptions struct {
	net_cc.NetOptions
}

func (c *client) Ping(payload uint64) (out uint64, err error) {
	ctx := c.PingOptions.PrepareContext()
	req := &proto.Ping_Request{Payload: payload}
	res := &proto.Ping_Response{}
	if res, err = c.grpcClient.Ping(ctx, req); err != nil {
		return
	}
	return res.Payload, nil
}
