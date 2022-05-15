package service

import (
	"crypto/tls"
	net_svc "github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

const (
	tcp4Network = "tcp4"
	tcp6Network = "tcp6"
	udsNetwork  = "unixpacket"
)

type endpoint struct {
	service        BaseServiceInterface
	lis            net.Listener
	grpcSrv        *grpc.Server
	grpcSrvOptions []grpc.ServerOption
}

func (e *endpoint) GrpcServer() *grpc.Server {
	return e.grpcSrv
}

type tcpEndpoint struct {
	endpoint
	hostname      string
	port          int
	httpTransport bool
	tlsCfg        *tls.Config
}

func (e *tcpEndpoint) Address() string {
	addr := e.hostname + ":" + strconv.Itoa(e.port)
	return addr
}

func (e *tcpEndpoint) listen(network string) error {
	if !e.httpTransport {
		lis, err := net.Listen(network, e.Address())
		if err != nil {
			return err
		}
		e.lis = lis
	} else {
		lis, err := tls.Listen(network, e.Address(), e.tlsCfg)
		if err != nil {
			return err
		}
		e.lis = lis
	}
	return nil
}

func (e *endpoint) serveInit() {
	e.grpcSrv = grpc.NewServer(e.grpcSrvOptions...)
	e.service.Service_Register(e.grpcSrv)
}

func (e *tcpEndpoint) tcpServe() error {
	e.serveInit()
	// Register go_srv.net.NetService
	net_svc.RegisterNetParcelServer(e.GrpcServer(), e.service.(NetworkServiceInterface).NetParcel())
	if !e.httpTransport {
		if err := e.grpcSrv.Serve(e.lis); err != nil {
			return err
		}
	} else {

	}
	return nil
}

func (e *localEndpoint) unixServe() error {
	e.serveInit()
	if err := e.grpcSrv.Serve(e.lis); err != nil {
		return err
	}
	return nil
}

//
// TCP 4 endpoint
//
type tcp4Endpoint struct {
	tcpEndpoint
}

func (e *tcp4Endpoint) Listen() error {
	return e.listen(tcp4Network)
}

//
// TCP 6 endpoint
//
type tcp6Endpoint struct {
	tcpEndpoint
	fallback *tcp4Endpoint
}

func (e *tcp6Endpoint) Listen() error {
	return e.listen(tcp6Network)
}

//
// Local endpoint
//
type localEndpoint struct {
	endpoint
	pathname string
}

func (e *localEndpoint) Address() string {
	return e.pathname
}
