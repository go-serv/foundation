package server

import (
	"context"
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
)

func (n *netParcel) Ping(ctx context.Context, req *proto.Ping_Request) (res *proto.Ping_Response, err error) {
	res = &proto.Ping_Response{}
	res.Payload = req.Payload
	return
}
