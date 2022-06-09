package ftp

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (FtpImpl) FtpInquiry(context.Context, *proto.Ftp_Inquiry_Request) (*proto.Ftp_Inquiry_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpInquiry not implemented")
}
