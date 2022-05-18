package server

import (
	"crypto/tls"
	rt "github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type endpoint struct {
	Srv                       serverInterface
	Lis                       net.Listener
	grpcSrv                   *grpc.Server
	GrpcSrvOptions            []grpc.ServerOption
	GrpcSrvUnaryInterceptors  []grpc.UnaryServerInterceptor
	GrpcSrvStreamInterceptors []grpc.StreamServerInterceptor
}

func (e *endpoint) WithServer(s serverInterface) {
	e.Srv = s
}

func (e *endpoint) GrpcServer() *grpc.Server {
	return e.grpcSrv
}

type tcpEndpoint struct {
	endpoint
	srv           service.NetworkServiceInterface
	hostname      string
	port          int
	httpTransport bool
	tlsCfg        *tls.Config
}

func (e *tcpEndpoint) serveInit() {
	// Build unary interceptors chain
	//e.GrpcSrvUnaryInterceptors = append(e.GrpcSrvUnaryInterceptors, server_md.PostUnaryInterceptor())
	//e.GrpcSrvUnaryInterceptors = append([]grpc.UnaryServerInterceptor{server_md.PreUnaryInterceptor()},
	//	e.GrpcSrvUnaryInterceptors...)
	//
	//e.GrpcSrvOptions = append(e.GrpcSrvOptions, grpc.ChainUnaryInterceptor(e.GrpcSrvUnaryInterceptors...))
	e.grpcSrv = grpc.NewServer(e.GrpcSrvOptions...)
	for _, svc := range rt.Runtime().NetworkServices() {
		svc.Service_Register(e.grpcSrv)
	}
}

func (e *tcpEndpoint) Address() string {
	addr := e.hostname + ":" + strconv.Itoa(e.port)
	return addr
}

func (e *tcpEndpoint) listen(network string) error {
	if !e.httpTransport {
		addr := e.Address()
		lis, err := net.Listen(network, addr)
		if err != nil {
			return err
		}
		e.Lis = lis
	} else {
		lis, err := tls.Listen(network, e.Address(), e.tlsCfg)
		if err != nil {
			return err
		}
		e.Lis = lis
	}
	return nil
}

func (e *tcpEndpoint) tcpServe() error {
	e.serveInit()
	if !e.httpTransport {
		if err := e.grpcSrv.Serve(e.Lis); err != nil {
			return err
		}
	} else {

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
	return e.listen(Tcp4Network)
}

//
// TCP 6 endpoint
//
type tcp6Endpoint struct {
	tcpEndpoint
	fallback *tcp4Endpoint
}

func (e *tcp6Endpoint) Listen() error {
	return e.listen(Tcp6Network)
}

func (e *endpoint) serveInit() {
	// Build unary interceptors chain
	//e.GrpcSrvUnaryInterceptors = append(e.GrpcSrvUnaryInterceptors, server_md.PostUnaryInterceptor())
	//e.GrpcSrvUnaryInterceptors = append([]grpc.UnaryServerInterceptor{server_md.PreUnaryInterceptor()},
	//	e.GrpcSrvUnaryInterceptors...)
	//
	//e.GrpcSrvOptions = append(e.GrpcSrvOptions, grpc.ChainUnaryInterceptor(e.GrpcSrvUnaryInterceptors...))
	//e.grpcSrv = grpc.NewServer(e.GrpcSrvOptions...)
	//e.service.Service_Register(e.grpcSrv)
}

//func (e *endpoint) netServeInit() {
//	for _, svc := range rt.Runtime().NetworkServices() {
//		svc.Service_Register(e.grpcSrv)
//	}
//}
//
//func (e *localEndpoint) unixServe() error {
//	e.serveInit()
//	if err := e.grpcSrv.Serve(e.Lis); err != nil {
//		return err
//	}
//	return nil
//}

//
// Local endpoint
//
//type localEndpoint struct {
//	endpoint
//	pathname string
//}
//
//func (e *localEndpoint) Address() string {
//	return e.pathname
//}
