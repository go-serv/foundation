package client

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
)

func (c *client) FtpInquiry(ctx context.Context, in *proto.Ftp_Inquiry_Request, opts ...grpc.CallOption) (*proto.Ftp_Inquiry_Response, error) {
	return c.stubs.FtpInquiry(ctx, in)
}
