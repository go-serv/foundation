//
// Implementation of NetParcel network service along with its runtime registration.

package net_parcel

import (
	"context"
	net_svc "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/grpc"
)

var Name = net_svc.NetParcel_ServiceDesc.ServiceName

type netParcel struct {
	service.NetworkServiceInterface
	net_svc.NetParcelServer
}

func (s *netParcel) Service_Register(srv *grpc.Server) {
	net_svc.RegisterNetParcelServer(srv, s)
}

func (s *netParcel) GetCryptoNonce(context.Context, *net_svc.CryptoNonce_Request) (*net_svc.CryptoNonce_Response, error) {
	return nil, nil
}
