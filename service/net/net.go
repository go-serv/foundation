package net

import (
	"github.com/mesh-master/foundation/pkg/z"
)

type X509PemPair struct {
	CertFile string
	KeyFile  string
}

type netService struct {
	z.ServiceInterface
	tlsEnabled bool
}

func (s *netService) WithTlsEnabled(flag bool) {
	s.tlsEnabled = flag
}

func (s *netService) IsTlsEnabled() bool {
	return s.tlsEnabled
}
