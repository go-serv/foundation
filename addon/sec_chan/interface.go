package sec_chan

import "github.com/go-serv/foundation/pkg/z"

type ServiceInterface interface {
	z.NetworkServiceInterface
}

type ServiceConfigInterface interface{}

type ClientInterface interface {
	z.NetworkClientInterface
	Create(lifetime int, nonceLen int) (pubKey []byte, nonce []byte, err error)
	Close() error
}

type ClientConfigInterface interface{}
