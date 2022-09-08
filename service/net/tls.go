package net

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
)

func (n *tcpEndpoint) WithNoTrustedPartiesTlsProfile(rootCertPemFile string, serverCertPairs []*X509PemPair) error {
	return n.loadCertificates(rootCertPemFile, serverCertPairs, tls.NoClientCert)
}

func (n *tcpEndpoint) WithTrustedPartiesTlsProfile(rootCertPemFile string, serverCertPairs []*X509PemPair) error {
	return n.loadCertificates(rootCertPemFile, serverCertPairs, tls.RequireAndVerifyClientCert)
}

func (ep *tcpEndpoint) loadCertificates(rootCertPemFile string, serverCertPairs []*X509PemPair, authType tls.ClientAuthType) (err error) {
	var (
		srvCert          tls.Certificate
		rootCertPemBytes []byte
	)

	rootCertPool := x509.NewCertPool()
	ep.tlsCfg = &tls.Config{
		// Assume that for the client/server authentication we will use a common root certificate.
		RootCAs:      rootCertPool,
		ClientCAs:    rootCertPool,
		Certificates: make([]tls.Certificate, 0),
		ClientAuth:   authType,
	}

	if len(rootCertPemFile) > 0 {
		if rootCertPemBytes, err = ioutil.ReadFile(rootCertPemFile); err != nil {
			return
		}
		ok := ep.tlsCfg.RootCAs.AppendCertsFromPEM(rootCertPemBytes)
		ok2 := ep.tlsCfg.ClientCAs.AppendCertsFromPEM(rootCertPemBytes)
		if !ok || !ok2 {
			return errors.New("loading CA root certificate failed")
		}
	}

	//
	for _, p := range serverCertPairs {
		if srvCert, err = tls.LoadX509KeyPair(p.CertFile, p.KeyFile); err != nil {
			return err
		}
		ep.tlsCfg.Certificates = append(ep.tlsCfg.Certificates, srvCert)
	}
	return
}
