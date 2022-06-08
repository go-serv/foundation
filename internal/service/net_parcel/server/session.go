package server

import (
	"context"
	"crypto/rand"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s serviceImpl) SecureSession(ctx context.Context, req *proto.Session_Request) (res *proto.Session_Response, err error) {
	var (
		nonce, encKey []byte
	)
	netCtx := ctx.(z.NetContextInterface)
	res = &proto.Session_Response{}
	// Create nonce
	if req.GetNonceLength() > 32 {
		return nil, status.Error(codes.InvalidArgument, "nonce length exceeds the maximum value of 32")
	}
	nonce = make([]byte, req.GetNonceLength())
	rand.Read(nonce)
	res.Nonce = nonce
	// Create session
	lifetime := uint16(req.GetLifetime())
	sess := session.NewSecureSession(lifetime, nonce, encKey)
	dic := netCtx.Response().Meta().Dictionary().(*net.HttpClientDictionary)
	dic.SessionId = sess.Id()
	return res, nil
}
