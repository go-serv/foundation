package client

import (
	i "github.com/go-serv/service/internal"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	"github.com/go-serv/service/internal/middleware/codec/cipher_msg"
	"github.com/go-serv/service/internal/runtime"
	net_service "github.com/go-serv/service/internal/service/net"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func newClient(e i.EndpointInterface) client {
	c := client{}
	c.endpoint = e
	c.insecure = true
	return c
}

func NewLocalClient(e i.EndpointInterface) *localClient {
	c := &localClient{newClient(e)}
	//runtime.Runtime().RegisterLocalClient(svcName, c)
	return c
}

func NewNetClient(svcName protoreflect.FullName, e i.EndpointInterface) *netClient {
	c := &netClient{}
	c.client = newClient(e)
	c.svc = net_service.NewNetworkService(svcName)
	// Set client codec
	codec := net_cc.NewOrRegistered(string(svcName))
	c.WithCodec(codec)
	cipher_msg.NetClientInit(c)
	//
	c.dialOpts = append(c.dialOpts,
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)),
		//grpc.WithChainUnaryInterceptor(c.mwGroup.UnaryClientInterceptor()),
	)
	runtime.Runtime().RegisterNetworkClient(svcName, c)
	return c
}
