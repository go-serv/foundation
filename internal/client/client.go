package client

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	mw_codec "github.com/go-serv/service/internal/grpc/mw_group/codec"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net"
)

type client struct {
	svcName      protoreflect.FullName
	codec        i.CodecInterface
	codecMwGroup i.CodecMiddlewareGroupInterface
	endpoint     i.EndpointInterface
	conn         net.Conn
	dialOpts     []grpc.DialOption
	mwGroup      i.MiddlewareGroupInterface
	insecure     bool
	ancillary.MethodMustBeImplemented
}

type localClient struct {
	client
}

type netClient struct {
	client
	svc i.NetworkServiceInterface
}

func (c *client) ServiceName() protoreflect.FullName {
	return c.svcName
}

func (c *client) Codec() i.CodecInterface {
	return c.codec
}

func (s *client) WithCodec(cc i.CodecInterface) {
	s.codec = cc
	s.codecMwGroup = mw_codec.NewCodecMiddlewareGroup(cc)
}

func (s *client) CodecMiddlewareGroup() i.CodecMiddlewareGroupInterface {
	return s.codecMwGroup
}

func (c *client) Endpoint() i.EndpointInterface {
	return c.endpoint
}

func (c *client) NewClient(cc grpc.ClientConnInterface) {
	c.MethodMustBeImplemented.Panic()
}

func (c *netClient) NetService() i.NetworkServiceInterface {
	return c.svc
}
