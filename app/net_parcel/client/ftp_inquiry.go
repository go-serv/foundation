package client

import (
	"context"
	proto "github.com/go-serv/foundation/internal/autogen/foundation"
	"google.golang.org/grpc"
)

func (c *client) FtpInquiry(ctx context.Context, in *proto.Ftp_Inquiry_Request, opts ...grpc.CallOption) (*proto.Ftp_Inquiry_Response, error) {
	return c.grpcClient.FtpInquiry(ctx, in)
}
