package client

import (
	"context"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/autogen/proto/net"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
)

var serviceName = proto.NetParcel_ServiceDesc.ServiceName

type client struct {
	i.NetworkClientInterface
	net.NetParcelClient
}

func (c *client) Client_NewClient(cc grpc.ClientConnInterface) {
	c.NetParcelClient = net.NewNetParcelClient(cc)
}

func (c *client) GetCryptoNonce(ctx context.Context, in *proto.CryptoNonce_Request, opts ...grpc.CallOption) (*proto.CryptoNonce_Response, error) {
	return c.NetParcelClient.GetCryptoNonce(ctx, in)
}
