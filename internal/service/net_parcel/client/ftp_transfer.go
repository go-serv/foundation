package client

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
)

func (c *client) FtpTransfer(ctx context.Context, in *proto.Ftp_FileChunk_Request, opts ...grpc.CallOption) (*proto.Ftp_FileChunk_Response, error) {
	return c.stubs.FtpTransfer(ctx, in)
}
