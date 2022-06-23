//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/service/net_parcel/server/ftp"
	"github.com/go-serv/service/pkg/y/netparcel"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/ancillary"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var Name = protoreflect.FullName(proto.NetParcel_ServiceDesc.ServiceName)

type serviceUnimpl struct {
	proto.UnimplementedNetParcelServer
}

type netParcel struct {
	z.NetworkServiceInterface
	ancillary.ArchiveOptions
	ftp.FtpImpl
	sessionImpl
	serviceUnimpl
}

func (s *netParcel) Register(srv *grpc.Server) {
	proto.RegisterNetParcelServer(srv, s)
}

func (s *netParcel) RegisterFtpPostActionHandler(fn netparcel.FtpPostActionHandlerFn, fileExt string) {
	s.FtpImpl.PostActions[fileExt] = fn
}
