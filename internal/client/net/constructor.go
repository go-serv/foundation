package net

import (
	"github.com/go-serv/service/internal/client"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	meta_net "github.com/go-serv/service/internal/grpc/meta/net"
	net_mw "github.com/go-serv/service/internal/grpc/middleware/net"
	"github.com/go-serv/service/internal/middleware/codec/cipher_msg"
	session_mw "github.com/go-serv/service/internal/middleware/net/session"
	"github.com/go-serv/service/internal/runtime"
	net_service "github.com/go-serv/service/internal/service/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewClient(svcName protoreflect.FullName, e z.EndpointInterface) *netClient {
	c := new(netClient)
	c.insecure = true
	c.ClientInterface = client.NewClient(svcName, e)
	c.svc = net_service.NewNetworkService(svcName)
	c.WithMeta(meta_net.NewMeta(nil))
	// Set client codec
	codec := net_cc.NewOrRegistered(string(svcName))
	c.WithCodec(codec)
	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	//
	cipher_msg.NetClientInit(c)
	c.newDefaultMiddleware()
	runtime.Runtime().RegisterNetworkClient(c)
	return c
}

func (c *netClient) newDefaultMiddleware() {
	mw := net_mw.NewMiddleware()
	mw.WithClient(c)
	c.WithMiddleware(mw)
	session_mw.ClientInit(mw)
}
