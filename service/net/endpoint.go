package net

import (
	"context"
	"crypto/tls"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strconv"
)

type tcpEndpoint struct {
	z.EndpointInterface
	hostname    string
	port        int
	tlsCfg      *tls.Config
	wrappedGrpc *grpcweb.WrappedGrpcServer
	webProxy    z.WebProxyInterface
}

func (ep *tcpEndpoint) Address() string {
	addr := ep.hostname + ":" + strconv.Itoa(ep.port)
	return addr
}

func (ep *tcpEndpoint) WithWebProxy(wp z.WebProxyInterface) {
	ep.webProxy = wp
}

func (ep *tcpEndpoint) IsSecure() bool {
	return ep.tlsCfg != nil
}

func (ep *tcpEndpoint) TlsConfig() *tls.Config {
	return ep.tlsCfg
}

func (ep *tcpEndpoint) TransportCredentials() credentials.TransportCredentials {
	if ep.tlsCfg == nil {
		return insecure.NewCredentials()
	} else {
		return credentials.NewTLS(ep.tlsCfg)
	}
}

func (ep *tcpEndpoint) ClientHandshake(ctx context.Context, s string, conn net.Conn) (net.Conn, credentials.AuthInfo, error) {
	return nil, nil, nil
}

func (ep *tcpEndpoint) listenAndServeNetwork(network string) (err error) {
	var (
		lis net.Listener
	)
	if ep.tlsCfg != nil {
		if lis, err = tls.Listen(network, ep.Address(), ep.tlsCfg); err != nil {
			return
		}
	} else {
		if lis, err = net.Listen(network, ep.Address()); err != nil {
			return
		}
	}
	//
	if ep.webProxy != nil {
		var httpSrv *http.Server
		if httpSrv, err = ep.webProxy.BuildHttpServer(ep.GrpcServer(), ep); err != nil {
			return
		}
		err = httpSrv.Serve(lis)
	} else {
		err = ep.GrpcServer().Serve(lis)
	}
	return
}

//
// TCP 4 endpoint
//
type tcp4Endpoint struct {
	tcpEndpoint
}

func (e *tcp4Endpoint) listenAndServe() error {
	return e.listenAndServeNetwork("tcp4")
}

//
// TCP 6 endpoint
//
type tcp6Endpoint struct {
	tcpEndpoint
	fallback *tcp4Endpoint
}

func (e *tcp6Endpoint) listenAndServe() error {
	return e.listenAndServeNetwork("tcp6")
}
