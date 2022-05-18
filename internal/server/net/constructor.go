package net

import (
	"github.com/go-serv/service/internal/server"
	_ "github.com/go-serv/service/internal/service/net_parcel" // this will enable the NetParcel service
)

func NewNetServer() *netServer {
	srv := new(netServer)
	srv.Server = server.NewBaseServer()
	return srv
}
