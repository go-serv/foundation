package client

import (
	proto "github.com/go-serv/foundation/internal/autogen/proto/net"
	client2 "github.com/go-serv/foundation/internal/client"
	"github.com/go-serv/foundation/pkg/ancillary/struc/copyable"
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
