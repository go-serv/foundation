package client

import (
	proto "github.com/go-serv/foundation/internal/autogen/proto/net"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

var (
	serviceName = proto.NetParcel_ServiceDesc.ServiceName
)

type client struct {
	z.NetworkClientInterface
	SecureSessionOptions
	FtpNewSessionOptions
	FtpTransferOptions
	stubs proto.NetParcelClient
}

func (c *client) OnConnect(cc grpc.ClientConnInterface) {
	c.stubs = proto.NewNetParcelClient(cc)
	c.MaxChunkSize = z.GrpcMaxMessageSize
	c.SecureSessionOptions.c = c
	c.FtpNewSessionOptions.c = c
	c.FtpTransferOptions.c = c
}
