package server

import (
	"context"
	"crypto/rand"
	"fmt"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const NonceMaxLength = 32

func (s serviceImpl) SecureSession(ctx context.Context, req *proto.Session_Request) (res *proto.Session_Response, err error) {
	var (
		nonce, encKey []byte
	)
	netCtx := ctx.(z.NetContextInterface)
	res = &proto.Session_Response{}
	// Create a nonce
	if req.GetNonceLength() > NonceMaxLength {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("nonce length exceeds the maximum value of %d", NonceMaxLength))
	}
	nonce = make([]byte, req.GetNonceLength(), NonceMaxLength)
	rand.Read(nonce)
	res.Nonce = nonce
	// Create a new session
	lifetime := uint16(req.GetLifetime())
	sess := session.NewSecureSession(lifetime, nonce, encKey)
	dic := netCtx.Response().Meta().Dictionary().(*net.HttpDictionary)
	dic.SessionId = sess.Id()
	return res, nil
}
