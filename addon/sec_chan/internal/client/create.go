package client

import (
	"github.com/go-serv/foundation/addon/sec_chan/x"
	"github.com/go-serv/foundation/internal/ancillary/crypto/aes"
	"github.com/go-serv/foundation/internal/ancillary/crypto/dh_key"
	proto "github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	grpc_client "github.com/go-serv/foundation/internal/client"
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/pkg/ancillary/struc/copyable"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
)

type SecureSessionOptions struct {
	copyable.Shallow
	grpc_client.NetOptions
	c *client
}

func (i impl) Create(in *proto.Create_Request, algoType x.KeyExchangeAlgoType) (res *proto.Create_Response, err error) {
	var (
		pubKeyExch  crypto.PubKeyExchangeInterface
		encKey      []byte
		blockCipher crypto.AEAD_CipherInterface
	)
	ctx := i.PrepareContext()

	// Public key exchange. Default to the Diffie-Hellman algorithm.
	switch algoType {
	case x.KeyExchangeDH:
		if pubKeyExch, err = dh_key.NewKeyExchange(); err != nil {
			return nil, err
		}
		in.KeyExchAlgo = &proto.Create_Request_Dh{Dh: &proto.Crypto_PubKeyExchange_DiffieHellman{PubKey: pubKeyExch.PublicKey()}}
	case x.KeyExchangeECDH:
		// TODO must be implemented
	case x.KeyExchangeRSA:
		// TODO must be implemented
	case x.KeyExchangePSK:
		in.KeyExchAlgo = &proto.Create_Request_Psk{}
	default:
		panic("secure channel: unimplemented key exchange algorithm")
	}

	if res, err = i.c.grpcClient.Create(ctx, in); err != nil {
		return
	}

	switch algoType {
	case x.KeyExchangeDH:
		if encKey, err = pubKeyExch.ComputeKey(res.GetPubKey()); err != nil {
			return nil, err
		}
	case x.KeyExchangePSK:
		var resolved any
		if resolved, err = runtime.Runtime().Resolve(x.PskResolverKey); err != nil {
			return
		}
		encKey = resolved.([]byte)
	}

	switch in.GetBlockCypher() {
	case proto.Crypto_AEADCipher_AES_GCM: // default
		if blockCipher, err = aes.NewCipher(encKey, res.GetNonce()); err != nil {
			return nil, err
		}
	}
	i.c.WithBlockCipher(blockCipher)
	return
}
