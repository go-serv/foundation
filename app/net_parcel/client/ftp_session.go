package client

import (
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	client2 "github.com/mesh-master/foundation/internal/grpc/client"
	"github.com/mesh-master/foundation/pkg/ancillary/struc/copyable"
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
	return f.c.grpcClient.FtpNewSession(ctx, in)
}
