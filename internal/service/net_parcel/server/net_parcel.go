//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	"context"
	"crypto/rand"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/service/net"
	"google.golang.org/grpc"
)

var Name = proto.NetParcel_ServiceDesc.ServiceName

type netParcel struct {
	net.NetworkServiceInterface
	proto.NetParcelServer
}

func (s *netParcel) Service_Register(srv *grpc.Server) {
	proto.RegisterNetParcelServer(srv, s)
}

func (s *netParcel) Service_OnNewSession(req request.RequestInterface) error {
	return nil
}

func (s *netParcel) GetCryptoNonce(ctx context.Context, req *proto.CryptoNonce_Request) (*proto.CryptoNonce_Response, error) {
	res := &proto.CryptoNonce_Response{}
	nonce := make([]byte, req.GetLen())
	rand.Read(nonce)
	res.Nonce = nonce
	return res, nil
}
