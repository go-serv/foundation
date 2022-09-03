package net

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func (n *netServer) loadCertificates(rootCertPemFile string, serverCertPairs []X509PemPair, authType tls.ClientAuthType) (err error) {
	var (
		srvCert          tls.Certificate
		rootCertPemBytes []byte
	)
	n.tlsCfg = &tls.Config{
		RootCAs:      &x509.CertPool{},
		Certificates: make([]tls.Certificate, 0),
		ClientAuth:   authType,
	}
	if rootCertPemBytes, err = ioutil.ReadFile(rootCertPemFile); err == nil {
		if ok := n.tlsCfg.RootCAs.AppendCertsFromPEM(rootCertPemBytes); !ok {
			return ErrRootCertificate
		}
	}
	//
	for _, p := range serverCertPairs {
		if srvCert, err = tls.LoadX509KeyPair(p.certFile, p.keyFile); err != nil {
			return err
		}
		n.tlsCfg.Certificates = append(n.tlsCfg.Certificates, srvCert)
	}
	return
}
