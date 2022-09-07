package net

import (
	"github.com/go-serv/foundation/pkg/z"
	"time"
)

type X509PemPair struct {
	CertFile string
	KeyFile  string
}

type WebProxyConfig struct {
	UseWebsocket          bool
	WsPingInterval        time.Duration
	WsReadLimit           int64
	PemCert               *X509PemPair
	AllowedOrigins        []string
	AllowedRequestHeaders []string
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
