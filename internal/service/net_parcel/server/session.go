package server

import (
	"context"
	"crypto/rand"
	"github.com/go-serv/service/internal/ancillary/crypto/dh_key"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	grpc_err "github.com/go-serv/service/internal/grpc/error"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/ancillary/crypto"
	"google.golang.org/grpc/codes"
)

const NonceMaxLength = 32

type sessionImpl struct {
}

func (s sessionImpl) diffieHellmanKeyExchange(req *proto.Session_Request) (serverPubKey []byte, encKey []byte, err error) {
	var (
		pubKeyExch   crypto.PubKeyExchangeInterface
		clientPubKey []byte
	)
	if pubKeyExch, err = dh_key.NewKeyExchange(); err != nil {
		return nil, nil, err
	}
	clientPubKey = req.GetKeyExchAlgo().(*proto.Session_Request_Dh).Dh.GetPubKey()
	if encKey, err = pubKeyExch.ComputeKey(clientPubKey); err != nil {
		return nil, nil, err
	}
	serverPubKey = pubKeyExch.PublicKey()
	return
}

func (s sessionImpl) SecureSession(ctx context.Context, req *proto.Session_Request) (res *proto.Session_Response, err error) {
	var (
		nonce, encKey, serverPubKey []byte
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
		return nil, grpc_err.New(z.ErrSeverityHigh, codes.FailedPrecondition, err.Error())
	}
	res.Nonce = nonce
	t := req.GetKeyExchAlgo()
	_ = t
	// Public key exchange.
	switch req.GetKeyExchAlgo().(type) {
	case *proto.Session_Request_Dh:
		if serverPubKey, encKey, err = s.diffieHellmanKeyExchange(req); err != nil {
			return nil, grpc_err.New(z.ErrSeverityHigh, codes.FailedPrecondition, err.Error())
		}
	case *proto.Session_Request_Ecdh:
		// TODO must be implemented
	case *proto.Session_Request_Rsa:
		// TODO must be implemented
	case *proto.Session_Request_Psk:
		// TODO must be implemented
	default:
		return nil, grpc_err.New(z.ErrSeverityLow, codes.FailedPrecondition, "public key exchange algo must be specified")
	}
	// Create a new session.
	lifetime := uint16(req.GetLifetime())
	sess := session.NewSecureSession(lifetime, nonce, encKey)
	dic := netCtx.Response().Meta().Dictionary().(*net.HttpDictionary)
	dic.SessionId = sess.Id()
	// Send server's public key to the client if necessary.
	if len(serverPubKey) > 0 {
		res.PubKey = serverPubKey
	}
	return res, nil
}
