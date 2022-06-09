package ftp

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (FtpImpl) FtpNewSession(context.Context, *proto.Ftp_NewSession_Request) (*proto.Ftp_NewSession_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpNewSession not implemented")
}
