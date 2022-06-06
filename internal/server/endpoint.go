package server

import (
	"crypto/tls"
	i "github.com/go-serv/service/internal"
	rt "github.com/go-serv/service/internal/runtime"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

type endpoint struct {
	srv                       i.ServerInterface
	lis                       net.Listener
	grpcSrv                   *grpc.Server
	GrpcSrvUnaryInterceptors  []grpc.UnaryServerInterceptor
	GrpcSrvStreamInterceptors []grpc.StreamServerInterceptor
}

func (e *endpoint) WithServer(s i.ServerInterface) {
	e.srv = s
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

func (e *endpoint) beforeServeInit() {
	//
	mw := e.srv.MiddlewareGroup()
	if mw != nil {
		ints := grpc.ChainUnaryInterceptor(mw.UnaryServerInterceptor())
		e.srv.AddGrpcServerOption(ints)
	}
	//
	e.grpcSrv = grpc.NewServer(e.srv.GrpcServerOptions()...)
	for _, svc := range rt.Runtime().RegisteredServices() {
		svc.Register(e.grpcSrv)
	}
}

func (e *tcpEndpoint) serveInit() {
	interceptors := grpc.ChainUnaryInterceptor(e.srv.MiddlewareGroup().UnaryServerInterceptor())
	e.srv.AddGrpcServerOption(interceptors)
	// Create a new gRPC server
	e.grpcSrv = grpc.NewServer(e.srv.GrpcServerOptions()...)
	// Register all network gRPC services
	for _, svc := range rt.Runtime().NetworkServices() {
		svc.Register(e.grpcSrv)
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

func (e *tcpEndpoint) tcpServe() error {
	e.beforeServeInit()
	if !e.httpTransport {
		if err := e.grpcSrv.Serve(e.lis); err != nil {
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

//
// Local endpoint
//
type localEndpoint struct {
	endpoint
	pathname string
}

func (e *localEndpoint) Listen() error {
	var err error
	var unixAddr *net.UnixAddr
	socketAddr := "@" + e.Address()
	unixAddr, err = net.ResolveUnixAddr(UnixDomainSocket, socketAddr)
	if err != nil {
		return err
	}
	e.lis, err = net.ListenUnix(UnixDomainSocket, unixAddr)
	if err != nil {
		return err
	}
	return nil
}

func (e *localEndpoint) unixServe() error {
	e.beforeServeInit()
	if err := e.grpcSrv.Serve(e.lis); err != nil {
		return err
	}
	return nil
}
