package client

import (
	"context"
	"github.com/go-serv/service/internal/autogen/proto/net"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var serviceName = protoreflect.FullName(proto.NetParcel_ServiceDesc.ServiceName)

type client struct {
	z.NetworkClientInterface
	stubs net.NetParcelClient
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.stubs = net.NewNetParcelClient(cc)
}

func (c *client) SecureSession(ctx context.Context, in *proto.Session_Request, opts ...grpc.CallOption) (*proto.Session_Response, error) {
	return c.stubs.SecureSession(ctx, in)
}
