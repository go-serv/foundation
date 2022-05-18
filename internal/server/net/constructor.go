package net

import (
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel" // this will enable the NetParcel service
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.Server = server.NewBaseServer()
	srv.mwGroup = defaultMiddlewareGroup(srv)
	return srv
}

func defaultMiddlewareGroup(srv *netServer) mw_group.MiddlewareGroupInterface {
	g := mw_group.NewMiddlewareGroup()
	g.AddItem(srv.sessionMwHandler)
	return g
}
