package server

import (
	"context"
	"crypto/rand"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	grpc_err "github.com/go-serv/service/internal/grpc/error"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
)

const NonceMaxLength = 32

type sessionImpl struct{}

func (s sessionImpl) SecureSession(ctx context.Context, req *proto.Session_Request) (res *proto.Session_Response, err error) {
	var (
		nonce, encKey []byte
	)
	netCtx := ctx.(z.NetContextInterface)
	res = &proto.Session_Response{}
	// Create a nonce with the given length.
	if req.GetNonceLength() > NonceMaxLength {
		return nil, grpc_err.New(
			z.ErrSeverityLow,
			codes.InvalidArgument,
			"nonce length exceeds the maximum value of %d",
			NonceMaxLength,
		)
	}
	nonce = make([]byte, req.GetNonceLength(), NonceMaxLength)
	if _, err = rand.Read(nonce); err != nil {
		return nil, grpc_err.New(z.ErrSeverityCritical, codes.FailedPrecondition, err.Error())
	}
	res.Nonce = nonce
	// Create a new session.
	lifetime := uint16(req.GetLifetime())
	sess := session.NewSecureSession(lifetime, nonce, encKey)
	dic := netCtx.Response().Meta().Dictionary().(*net.HttpDictionary)
	dic.SessionId = sess.Id()
	return res, nil
}
