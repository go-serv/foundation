//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	"context"
	"crypto/rand"
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

func (s *netParcel) GetCryptoNonce(ctx context.Context, req *net_svc.CryptoNonce_Request) (*net_svc.CryptoNonce_Response, error) {
	res := &net_svc.CryptoNonce_Response{}
	nonce := make([]byte, req.GetLen())
	rand.Read(nonce)
	res.Nonce = nonce
	return res, nil
}
