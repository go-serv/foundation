package net

import (
	"crypto/tls"
	"errors"
	"github.com/go-serv/service/pkg/z"
)

var (
	ErrRootCertificate = errors.New("failed to load CA root certificate")
)

type netServer struct {
	z.ServerInterface
	tlsCfg         *tls.Config
	enableWebProxy bool
}

type X509PemPair struct {
	certFile string
	keyFile  string
}

func (n *netServer) WithNoTrustedPartiesTlsProfile(rootCertPemFile string, serverCertPairs []X509PemPair) error {
	return n.loadCertificates(rootCertPemFile, serverCertPairs, tls.NoClientCert)
}

func (n *netServer) WithTrustedPartiesTlsProfile(rootCertPemFile string, serverCertPairs []X509PemPair) error {
	return n.loadCertificates(rootCertPemFile, serverCertPairs, tls.RequireAndVerifyClientCert)
}

func (srv *netServer) Resolver() z.NetworkServerResolverInterface {
	return nil
}

func (srv *netServer) WithResolver(resolver z.NetworkServerResolverInterface) {

}
