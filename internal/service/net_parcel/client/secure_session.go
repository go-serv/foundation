package client

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
)

func (c *client) SecureSession(ctx context.Context, in *proto.Session_Request, opts ...grpc.CallOption) (res *proto.Session_Response, err error) {
	res, err = c.stubs.SecureSession(ctx, in)
	if err != nil {
		return
	}
	netCtx := (ctx).(z.NetContextInterface)
	dic := netCtx.Response().Meta().Dictionary().(*net.HttpDictionary)
	_ = dic
	return
}
