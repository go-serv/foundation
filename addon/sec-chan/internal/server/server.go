//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	proto "github.com/go-serv/foundation/internal/autogen/sec_chan_mw"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/ancillary"
	"google.golang.org/grpc"
)

var Name = proto.SecureChannel_ServiceDesc.ServiceName

type serviceUnimpl struct {
	proto.UnimplementedSecureChannelServer
}

type secChanServer struct {
	z.NetworkServiceInterface
	ancillary.ArchiveOptions
	sessionImpl
	serviceUnimpl
}

type ConfigInterface interface {
	z.ServiceCfgInterface
}

func (s *secChanServer) Register(srv *grpc.Server) {
	proto.RegisterSecureChannelServer(srv, s)
}
