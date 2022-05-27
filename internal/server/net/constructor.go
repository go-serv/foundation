package net

import (
	i "github.com/go-serv/service/internal"
	mw_cipher "github.com/go-serv/service/internal/grpc/middleware/cipher_msg"
	mw "github.com/go-serv/service/internal/grpc/middleware/mw_group/net"
	mw_session "github.com/go-serv/service/internal/grpc/middleware/net_session"
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel/server" // this will enable the NetParcel service
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.Server = server.NewBaseServer()
	srv.Server.WithMiddlewareGroup(srv.defaultMiddlewareGroup())
	return srv
}

func (srv *netServer) defaultMiddlewareGroup() i.MiddlewareGroupInterface {
	g := mw.NewMiddlewareGroup(srv)
	g.AddItem(mw_session.NewNetSessionHandler())
	g.AddItem(mw_cipher.NewNetCipherServerHandler())
	return g
}
