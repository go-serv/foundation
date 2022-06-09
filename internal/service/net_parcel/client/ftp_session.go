package client

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
)

func (c *client) FtpNewSession(ctx context.Context, in *proto.Ftp_NewSession_Request, opts ...grpc.CallOption) (*proto.Ftp_NewSession_Response, error) {
	return c.stubs.FtpNewSession(ctx, in)
}
