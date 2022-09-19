package net

import (
	net_cc "github.com/go-serv/foundation/internal/grpc/codec/net"
	net_mw "github.com/go-serv/foundation/internal/grpc/middleware/net"
	enc_mw "github.com/go-serv/foundation/internal/middleware/net/enc_msg"
	session_mw "github.com/go-serv/foundation/internal/middleware/net/session"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
)

func defaultMiddlewareGroup() z.MiddlewareInterface {
	g := net_mw.NewMiddleware()
	session_mw.ServerInit(g)
	enc_mw.ServerInit(g)
	return g
}

func NewNetworkService(name string, cfg z.ServiceCfgInterface, endpoints []z.EndpointInterface) *netService {
	s := &netService{}
	s.ServiceInterface = service.NewService(name, cfg, endpoints)
	cc := net_cc.NewOrRegistered(name)
	s.ServiceInterface.WithCodec(cc)
	s.ServiceInterface.WithMiddlewareGroup(defaultMiddlewareGroup())
	return s
}

func newTcpEndpoint(hostname string, port int) tcpEndpoint {
	ep := tcpEndpoint{}
	ep.EndpointInterface = service.NewEndpoint()
	ep.hostname = hostname
	ep.port = port
	return ep
}

func NewTcp4Endpoint(hostname string, port int) *tcp4Endpoint {
	e := &tcp4Endpoint{newTcpEndpoint(hostname, port)}
	return e
}

func NewTcp6Endpoint(hostname string, port int) *tcp6Endpoint {
	e := &tcp6Endpoint{newTcpEndpoint(hostname, port), nil}
	return e
}

//func NewConfig(webProxy *WebProxyConfig) *cfg {
//	cfg := new(cfg)
//	cfg.WebProxyConfig = webProxy
//	return cfg
//}
