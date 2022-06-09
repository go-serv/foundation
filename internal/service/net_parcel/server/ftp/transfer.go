package ftp

import (
	"context"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (FtpImpl) FtpTransfer(context.Context, *proto.Ftp_FileChunk_Request) (*proto.Ftp_FileChunk_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FtpTransfer not implemented")
}
