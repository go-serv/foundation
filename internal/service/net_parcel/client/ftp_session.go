package client

import (
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	client2 "github.com/go-serv/service/internal/grpc/client"
)

type FtpNewSessionOptions struct {
	copyable.Shallow
	client2.NetOptions
	c        *client
	Lifetime *uint32
}

func (f FtpNewSessionOptions) FtpNewSession(in *proto.Ftp_NewSession_Request) (*proto.Ftp_NewSession_Response, error) {
	ctx := f.PrepareContext()
	in.Lifetime = f.Lifetime
	return f.c.stubs.FtpNewSession(ctx, in)
}
