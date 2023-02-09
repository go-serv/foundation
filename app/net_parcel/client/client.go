package client

import (
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/pkg/ancillary/struc/copyable"
	"github.com/mesh-master/foundation/pkg/z"
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
	PingOptions
	grpcClient proto.NetParcelClient
}

func (c *client) WithOptions(opts interface{}) {
	switch opts.(type) {
	case PingOptions:
		src := opts.(PingOptions)
		copyable.Shallow{}.Merge(c.PingOptions, src)
	}
}

func (c *client) OnConnect(cc grpc.ClientConnInterface) {
	c.grpcClient = proto.NewNetParcelClient(cc, "proto")
	c.MaxChunkSize = z.GrpcMaxMessageSize
	c.SecureSessionOptions.c = c
	c.FtpNewSessionOptions.c = c
	c.FtpTransferOptions.c = c
}
