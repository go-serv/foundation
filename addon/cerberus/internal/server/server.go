package server

import (
	proto "github.com/go-serv/foundation/internal/autogen/net/cerberus"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
)

var Name = proto.Cerberus_ServiceDesc.ServiceName

type serviceUnimpl struct {
	proto.UnimplementedCerberusServer
}

type impl struct{}

type cerberusService struct {
	z.NetworkServiceInterface
	impl
	serviceUnimpl
}

type ConfigInterface interface {
	z.ServiceCfgInterface
}

func (s *cerberusService) Register(srv *grpc.Server) {
	proto.RegisterCerberusServer(srv, s)
}
