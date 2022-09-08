package main

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/app"
	"github.com/go-serv/foundation/app/net_parcel/server"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/service/net"
	src "go-server-tests-endpoints"
	"log"
	"os"
)

var (
	rootCaCertFile string
	srvCertPemFile string
	srvCertKeyFile string
)

func createWebProxyCfg() *net.WebProxyConfig {
	proxyCfg := &net.WebProxyConfig{}
	proxyCfg.PemCert = &net.X509PemPair{
		os.Getenv(src.EnvCertServerPemFile),
		os.Getenv(src.EnvCertServerKeyFile),
	}
	return proxyCfg
}

func main() {
	var (
		err error
	)
	srvApp := app.NewApp()
	certs := make([]*net.X509PemPair, 1)

	// Read env variables, must be set in docker-compose file.
	rootCaCertFile = os.Getenv(src.EnvCertRootCaPemFile)
	if len(rootCaCertFile) == 0 {
		panic("Root CA certificate file, env variable not set")
	}
	srvCertPemFile = os.Getenv(src.EnvCertServerPemFile)
	srvCertKeyFile = os.Getenv(src.EnvCertServerKeyFile)
	if len(srvCertPemFile) == 0 || len(srvCertKeyFile) == 0 {
		panic("Server certificate file, env variable not set")
	}
	srvPemPair := &net.X509PemPair{srvCertPemFile, srvCertKeyFile}
	certs[0] = srvPemPair

	// Plain TCP connection with no TLS.
	eps := make([]z.EndpointInterface, 0)
	unsafeEp := net.NewTcp4Endpoint(src.ServerAddr, src.UnsafePort)
	eps = append(eps, unsafeEp)

	// Require an TLS authentication.
	trustedEp := net.NewTcp4Endpoint(src.ServerAddr, src.TlsTrustedPartiesPort)
	if err = trustedEp.WithTrustedPartiesTlsProfile(rootCaCertFile, certs); err != nil {
		panic(err)
	}
	eps = append(eps, trustedEp)

	// Do not authenticate the clients.
	noTrustEp := net.NewTcp4Endpoint(src.ServerAddr, src.TlsAnyPort)
	if err = noTrustEp.WithNoTrustedPartiesTlsProfile(rootCaCertFile, []*net.X509PemPair{srvPemPair}); err != nil {
		panic(err)
	}
	eps = append(eps, noTrustEp)

	// gRPC web proxy
	proxyEp := net.NewTcp4Endpoint(src.ServerAddr, src.WebProxyPort)
	proxyEp.WithWebProxy(createWebProxyCfg())
	eps = append(eps, proxyEp)

	// Creates NetParcel service and start the application.
	netParcel := server.NewNetParcel(eps, nil)
	srvApp.AddService(netParcel)
	srvApp.Start()
	if srvApp.Job().GetState() != job.Done {
		_, reason := srvApp.Job().GetInterruptedBy()
		log.Fatalf("failed with %v\n", reason)
	}
}
