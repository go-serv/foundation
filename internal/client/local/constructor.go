package local

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/client"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	mw_shmem "github.com/go-serv/service/internal/middleware/codec/shm_ipc"
	"github.com/go-serv/service/internal/runtime"
	loc_service "github.com/go-serv/service/internal/service/local"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewClient(svcName protoreflect.FullName, e i.EndpointInterface) *localClient {
	c := new(localClient)
	c.ClientInterface = client.NewClient(svcName, e)
	c.svc = loc_service.NewService(svcName)
	// Set client codec
	codec := net_cc.NewOrRegistered(string(svcName))
	c.WithCodec(codec)
	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	// Local client middlewares
	mw_shmem.ClientInit(c)
	//
	rt := runtime.Runtime()
	rt.RegisterLocalClient(c)
	return c
}
