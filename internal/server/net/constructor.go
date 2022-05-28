package net

import (
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel/server" // this will enable the NetParcel service
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.ServerInterface = server.NewServer()
	//srv.server.WithMiddlewareGroup(srv.defaultMiddlewareGroup())
	return srv
}

//func (srv *netServer) defaultMiddlewareGroup() i.MiddlewareGroupInterface {
//	g := mw.NewMiddlewareGroup(srv)
//	//g.AddItem(mw_session.NewNetSessionHandler())
//	//g.AddItem(mw_cipher.NewNetCipherServerHandler())
//	return g
//}
