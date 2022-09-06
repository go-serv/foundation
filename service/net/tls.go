package net

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func (n *tcpEndpoint) WithNoTrustedPartiesTlsProfile(rootCertPemFile string, serverCertPairs []X509PemPair) error {
	return n.loadCertificates(rootCertPemFile, serverCertPairs, tls.NoClientCert)
}

func (n *tcpEndpoint) WithTrustedPartiesTlsProfile(rootCertPemFile string, serverCertPairs []X509PemPair) error {
	return n.loadCertificates(rootCertPemFile, serverCertPairs, tls.RequireAndVerifyClientCert)
}

func (ep *tcpEndpoint) loadCertificates(rootCertPemFile string, serverCertPairs []X509PemPair, authType tls.ClientAuthType) (err error) {
	var (
		srvCert          tls.Certificate
		rootCertPemBytes []byte
	)
	ep.tlsCfg = &tls.Config{
		RootCAs:      &x509.CertPool{},
		Certificates: make([]tls.Certificate, 0),
		ClientAuth:   authType,
	}
	if rootCertPemBytes, err = ioutil.ReadFile(rootCertPemFile); err == nil {
		if ok := ep.tlsCfg.RootCAs.AppendCertsFromPEM(rootCertPemBytes); !ok {
			// todo: return an error
			return
		}
	}
	//
	for _, p := range serverCertPairs {
		if srvCert, err = tls.LoadX509KeyPair(p.certFile, p.keyFile); err != nil {
			return err
		}
		ep.tlsCfg.Certificates = append(ep.tlsCfg.Certificates, srvCert)
	}
	return
}
