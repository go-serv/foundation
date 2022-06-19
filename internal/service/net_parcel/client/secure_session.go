package client

import (
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/call"
)

type SecureSessionOptions struct {
	copyable.Shallow
	call.NetOptions
	c *client
}

func (s SecureSessionOptions) SecureSession(in *proto.Session_Request) (res *proto.Session_Response, err error) {
	ctx := s.PrepareContext()
	res, err = s.c.stubs.SecureSession(ctx, in)
	return
}
