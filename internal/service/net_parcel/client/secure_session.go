package client

import (
	"github.com/go-serv/service/internal/ancillary/crypto/aes"
	"github.com/go-serv/service/internal/ancillary/crypto/dh_key"
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	proto "github.com/go-serv/service/internal/autogen/proto/net"
	grpc_client "github.com/go-serv/service/internal/grpc/client"
	"github.com/go-serv/service/pkg/z/ancillary/crypto"
)

type SecureSessionOptions struct {
	copyable.Shallow
	grpc_client.NetOptions
	c *client
}

func (s SecureSessionOptions) SecureSession(in *proto.Session_Request) (res *proto.Session_Response, err error) {
	var (
		pubKeyExch  crypto.PubKeyExchangeInterface
		encKey      []byte
		blockCipher crypto.AEAD_CypherInterface
	)
	ctx := s.PrepareContext()
	// Public key exchange. Default to the Diffie-Hellman algorithm.
	switch in.GetKeyExchAlgo().(type) {
	case *proto.Session_Request_Dh:
		if pubKeyExch, err = dh_key.NewKeyExchange(); err != nil {
			return nil, err
		}
		in.KeyExchAlgo = &proto.Session_Request_Dh{Dh: &proto.Crypto_PubKeyExchange_DiffieHellman{PubKey: pubKeyExch.PublicKey()}}
	case *proto.Session_Request_Ecdh:
		// TODO must be implemented
	case *proto.Session_Request_Rsa:
		// TODO must be implemented
	case *proto.Session_Request_Psk:
		// TODO must be implemented
	default:
		if pubKeyExch, err = dh_key.NewKeyExchange(); err != nil {
			return nil, err
		}
		in.KeyExchAlgo = &proto.Session_Request_Dh{Dh: &proto.Crypto_PubKeyExchange_DiffieHellman{PubKey: pubKeyExch.PublicKey()}}
	}
	if res, err = s.c.stubs.SecureSession(ctx, in); err != nil {
		return
	}
	if pubKeyExch != nil {
		if encKey, err = pubKeyExch.ComputeKey(res.GetPubKey()); err != nil {
			return nil, err
		}
	}
	switch in.GetBlockCypher() {
	case proto.Crypto_AEADCipher_AES_GCM: // default
		if blockCipher, err = aes.NewCipher(encKey, res.GetNonce()); err != nil {
			return nil, err
		}
	}
	s.c.WithBlockCipher(blockCipher)
	return
}
