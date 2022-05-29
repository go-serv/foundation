package local

import (
	"github.com/go-serv/service/internal/server"
)

func NewServer() *localServer {
	srv := new(localServer)
	srv.ServerInterface = server.NewServer()
	return srv
}
