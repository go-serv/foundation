package server

import (
	"context"
	"crypto/rand"
	"github.com/go-serv/foundation/addon/sec_chan/x"
	"github.com/go-serv/foundation/internal/ancillary/crypto/aes"
	"github.com/go-serv/foundation/internal/ancillary/crypto/dh_key"
	proto "github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	grpc_err "github.com/go-serv/foundation/internal/grpc/error"
	"github.com/go-serv/foundation/internal/grpc/meta/net"
	"github.com/go-serv/foundation/internal/grpc/session"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
	"google.golang.org/grpc/codes"
)

const NonceMaxLength = 32

func (s impl) diffieHellmanKeyExchange(req *proto.Create_Request) (serverPubKey []byte, encKey []byte, err error) {
	var (
		pubKeyExch   crypto.PubKeyExchangeInterface
		clientPubKey []byte
	)
	if pubKeyExch, err = dh_key.NewKeyExchange(); err != nil {
		return
	}

	clientPubKey = req.GetKeyExchAlgo().(*proto.Create_Request_Dh).Dh.GetPubKey()
	if encKey, err = pubKeyExch.ComputeKey(clientPubKey); err != nil {
		return
	}

	serverPubKey = pubKeyExch.PublicKey()
	return
}

func (s impl) Create(ctx context.Context, req *proto.Create_Request) (res *proto.Create_Response, err error) {
	var (
		nonce, encKey, serverPubKey []byte
		cipher                      crypto.AEAD_CipherInterface
	)
	netCtx := ctx.(z.NetServerContextInterface)
	res = &proto.Create_Response{}

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

	// Public key exchange.
	switch req.GetKeyExchAlgo().(type) {
	case *proto.Create_Request_Dh:
		if serverPubKey, encKey, err = s.diffieHellmanKeyExchange(req); err != nil {
			return nil, grpc_err.New(z.ErrSeverityHigh, codes.FailedPrecondition, err.Error())
		}
	case *proto.Create_Request_Ecdh:
		// TODO must be implemented
	case *proto.Create_Request_Rsa:
		// TODO must be implemented
	case *proto.Create_Request_Psk: // Pre-shared key
		var resolved any
		if resolved, err = runtime.Runtime().Resolve(x.PskResolverKey); err != nil {
			return
		}
		encKey = resolved.([]byte)
	default:
		return nil, grpc_err.New(z.ErrSeverityLow, codes.FailedPrecondition, "public key exchange algo must be specified")
	}

	// Create a new session.
	lifetime := uint16(req.GetLifetime())
	sess := session.NewSession(lifetime)
	dic := netCtx.Response().Meta().Dictionary().(*net.HttpDictionary)
	dic.SessionId = sess.Id()
	netCtx.WithSession(sess)
	// Send back server's public key to the client if necessary.
	if len(serverPubKey) > 0 {
		res.PubKey = serverPubKey
	}

	// Init block cipher.
	switch req.GetBlockCypher() {
	case proto.Crypto_AEADCipher_AES_GCM:
		if cipher, err = aes.NewCipher(encKey, nonce); err != nil {
			return
		}
		sess.WithBlockCipher(cipher)
	}
	return res, nil
}
