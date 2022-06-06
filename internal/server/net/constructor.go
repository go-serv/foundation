package net

import (
	i "github.com/go-serv/service/internal"
	net_mw "github.com/go-serv/service/internal/grpc/mw_group/net"
	session_mw "github.com/go-serv/service/internal/middleware/net/session"
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel/server" // this will enable the NetParcel service
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.ServerInterface = server.NewServer()
	srv.ServerInterface.WithMiddlewareGroup(srv.defaultMiddlewareGroup())
	return srv
}

func (srv *netServer) defaultMiddlewareGroup() i.MiddlewareGroupInterface {
	g := net_mw.NewMiddlewareGroup()
	session_mw.ServerInit(g)
	return g
}
