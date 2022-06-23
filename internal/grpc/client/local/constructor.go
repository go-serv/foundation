package local

import (
	"github.com/go-serv/service/internal/grpc/client"
	local_cc "github.com/go-serv/service/internal/grpc/codec/local"
	mw_shmem "github.com/go-serv/service/internal/middleware/codec/shm_ipc"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewClient(svcName protoreflect.FullName, e z.EndpointInterface) *localClient {
	c := new(localClient)
	c.ClientInterface = client.NewClient(svcName, e)
	//c.svc = loc_service.NewService(svcName)
	// Set client codec
	codec := local_cc.NewOrRegistered(string(svcName))
	c.WithCodec(codec)
	c.WithDialOption(grpc.WithDefaultCallOptions(grpc.ForceCodec(codec)))
	// Local client middlewares
	mw_shmem.ClientInit(c)
	//
	rt := runtime.Runtime()
	rt.RegisterLocalClient(c)
	return c
}
