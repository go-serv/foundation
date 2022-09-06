package client

import (
	"context"
	i "github.com/go-serv/foundation/internal"
	"github.com/go-serv/foundation/internal/autogen/proto/local"
	proto "github.com/go-serv/foundation/internal/autogen/proto/local"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var serviceName = protoreflect.FullName(proto.Sample_ServiceDesc.ServiceName)

type client struct {
	i.LocalClientInterface
	stubs local.SampleClient
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.stubs = local.NewSampleClient(cc)
}

func (c *client) DoLargeRequest(ctx context.Context, in *proto.LargeRequest_Request, opts ...grpc.CallOption) (*proto.LargeRequest_Response, error) {
	return c.stubs.DoLargeRequest(ctx, in)
}

func (c *client) DoLargeRequestIpc(ctx context.Context, in *proto.LargeRequestIpc_Request, opts ...grpc.CallOption) (*proto.LargeRequestIpc_Response, error) {
	return c.stubs.DoLargeRequestIpc(ctx, in)
}
