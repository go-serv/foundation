package sec_chan

import "github.com/go-serv/foundation/pkg/z"

type ServiceInterface interface {
	z.NetworkServiceInterface
}

type ServiceConfigInterface interface{}

type ClientInterface interface {
	z.NetworkClientInterface
	NewSession(lifetime int, nonceLen int) (pubKey []byte, nonce []byte, err error)
}

type ClientConfigInterface interface{}
