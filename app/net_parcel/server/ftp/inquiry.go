package ftp

import (
	"context"
	proto "github.com/go-serv/foundation/internal/autogen/foundation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (FtpImpl) FtpInquiry(ctx context.Context, req *proto.Ftp_Inquiry_Request) (res *proto.Ftp_Inquiry_Response, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpInquiry not implemented")
}
