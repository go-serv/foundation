package server

import (
	"context"
	proto "github.com/go-serv/foundation/internal/autogen/proto/net"
)

func (n *netParcel) Ping(ctx context.Context, req *proto.Ping_Request) (res *proto.Ping_Response, err error) {
	res = &proto.Ping_Response{}
	res.Payload = req.Payload
	return
}
