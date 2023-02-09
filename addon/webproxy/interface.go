package sec_chan

import (
	"github.com/mesh-master/foundation/addon/sec_chan/x"
	"github.com/mesh-master/foundation/pkg/z"
)

type ServiceInterface interface {
	z.NetworkServiceInterface
}

type ServiceConfigInterface interface{}

type ClientInterface interface {
	z.NetworkClientInterface
	Create(lifetime int, nonceLen int, algoType x.KeyExchangeAlgoType) (pubKey []byte, nonce []byte, err error)
	Close() error
}

type ClientConfigInterface interface{}
