package net

import (
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	md_server "github.com/go-serv/service/internal/grpc/middleware/server"
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel/server" // this will enable the NetParcel service
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.Server = server.NewBaseServer()
	srv.Server.WithMiddlewareGroup(defaultMiddlewareGroup())
	return srv
}

func defaultMiddlewareGroup() mw_group.MiddlewareGroupInterface {
	g := mw_group.NewMiddlewareGroup()
	g.AddItem(md_server.NewNetSessionHandler())
	return g
}
