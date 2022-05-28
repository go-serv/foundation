package pkg

import (
	i "github.com/go-serv/service/internal"
	net_srv "github.com/go-serv/service/internal/server/net"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/grpc"
)

// All method names are prefixed with Service_ to avoid name conflicts with the method names of a GRPC service.
type serverInterface interface {
	Service_Name(bool) string
	Service_Register(srv *grpc.Server)
	AddEndpoint(endpoint i.EndpointInterface)
	Start()
	Stop()
	State() service.State
	// Adds a new wrapper to the wrapper chain
	// AddGrpcMessageWrapper(GrpcMessageWrapperFn)
}

type NetworkServerApi interface {
	i.NetworkServerInterface
}

type LocalServerApi interface {
}

func NewNetServer() NetworkServerApi {
	return net_srv.NewNetServer()
}
