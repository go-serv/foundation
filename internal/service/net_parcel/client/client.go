package client

import (
	"context"
	"github.com/go-serv/service/internal/autogen/proto/net"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	net_client "github.com/go-serv/service/internal/client"
	"google.golang.org/grpc"
)

type client struct {
	net_client.NetworkClientInterface
	net.NetParcelClient
}

func (c *client) Client_NewClient(cc grpc.ClientConnInterface) {
	c.NetParcelClient = net.NewNetParcelClient(cc)
}

func (c *client) GetCryptoNonce(ctx context.Context, in *proto.CryptoNonce_Request, opts ...grpc.CallOption) (*proto.CryptoNonce_Response, error) {
	return c.NetParcelClient.GetCryptoNonce(ctx, in)
}
