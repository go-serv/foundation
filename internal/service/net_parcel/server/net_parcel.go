//
// Implementation of NetParcel network service along with its runtime registration.

package server

import (
	"context"
	"crypto/rand"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var Name = protoreflect.FullName(proto.NetParcel_ServiceDesc.ServiceName)

type serviceImpl struct {
	proto.UnimplementedNetParcelServer
}

type netParcel struct {
	z.NetworkServiceInterface
	impl serviceImpl
}

func (s *netParcel) Register(srv *grpc.Server) {
	proto.RegisterNetParcelServer(srv, s.impl)
}

func (s *netParcel) Service_OnNewSession(req z.RequestInterface) error {
	return nil
}

func (s serviceImpl) GetCryptoNonce(ctx context.Context, req *proto.CryptoNonce_Request) (*proto.CryptoNonce_Response, error) {
	netCtx := ctx.(z.NetContextInterface)
	res := &proto.CryptoNonce_Response{}
	nonce := make([]byte, req.GetLen())
	rand.Read(nonce)
	res.Nonce = nonce
	r := netCtx.Request()
	_ = r
	return res, nil
}
