package net

import (
	_ "github.com/go-serv/service/internal/grpc/codec/net"
	md_cipher "github.com/go-serv/service/internal/grpc/middleware/cipher_msg"
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	md_session "github.com/go-serv/service/internal/grpc/middleware/net_session"
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
	g.AddItem(md_session.NewNetSessionHandler())
	g.AddItem(md_cipher.NewNetCipherServerHandler())
	return g
}
