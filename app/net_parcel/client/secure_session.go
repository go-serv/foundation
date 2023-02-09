package client

import (
	"github.com/mesh-master/foundation/internal/ancillary/crypto/aes"
	"github.com/mesh-master/foundation/internal/ancillary/crypto/dh_key"
	proto "github.com/mesh-master/foundation/internal/autogen/foundation"
	grpc_client "github.com/mesh-master/foundation/internal/grpc/client"
	"github.com/mesh-master/foundation/pkg/ancillary/struc/copyable"
	"github.com/mesh-master/foundation/pkg/z/ancillary/crypto"
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
		blockCipher crypto.AEAD_CipherInterface
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
	if res, err = s.c.grpcClient.SecureSession(ctx, in); err != nil {
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
