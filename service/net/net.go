package net

import (
	"github.com/go-serv/foundation/pkg/z"
)

type netService struct {
	z.ServiceInterface
	tenantId z.TenantId
	encKey   []byte
}

type X509PemPair struct {
	certFile string
	keyFile  string
}

//func (b *netService) EncriptionKey() []byte {
//	//TODO implement me
//	return []byte("secret")
//}
//
//func (b *netService) WithEncriptionKey(key []byte) {
//	b.encKey = key
//}
