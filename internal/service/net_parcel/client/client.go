package client

import (
	"context"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/autogen/proto/net"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var serviceName = protoreflect.FullName(proto.NetParcel_ServiceDesc.ServiceName)

type client struct {
	i.NetworkClientInterface
	stubs net.NetParcelClient
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.stubs = net.NewNetParcelClient(cc)
}

func (c *client) GetCryptoNonce(ctx context.Context, in *proto.CryptoNonce_Request, opts ...grpc.CallOption) (*proto.CryptoNonce_Response, error) {
	return c.stubs.GetCryptoNonce(ctx, in)
}
