package net

import (
	"github.com/go-serv/foundation/pkg/z"
)

type X509PemPair struct {
	CertFile string
	KeyFile  string
}

type netService struct {
	z.ServiceInterface
	tenantId z.TenantId
	encKey   []byte
}

//func (b *netService) EncriptionKey() []byte {
//	//TODO implement me
//	return []byte("secret")
//}
//
//func (b *netService) WithEncriptionKey(key []byte) {
//	b.encKey = key
//}
