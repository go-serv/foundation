package client

import (
	proto "github.com/go-serv/foundation/internal/autogen/proto/net"
	net_cc "github.com/go-serv/foundation/internal/client"
)

type PingOptions struct {
	net_cc.NetOptions
}

func (c *client) Ping(payload uint64) (out uint64, err error) {
	ctx := c.PingOptions.PrepareContext()
	req := &proto.Ping_Request{Payload: payload}
	res := &proto.Ping_Response{}
	if res, err = c.stubs.Ping(ctx, req); err != nil {
		return
	}
	return res.Payload, nil
}
