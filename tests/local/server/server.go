package main

import (
	"github.com/go-serv/foundation/internal/runtime"
	"github.com/go-serv/foundation/internal/server"
	"github.com/go-serv/foundation/internal/server/local"
	_ "github.com/go-serv/foundation/tests/local/server/internal"
)

func main() {
	svc := runtime.Runtime().LocalService()
	e := server.NewLocalEndpoint(svc)
	srv := local.NewServer()
	srv.AddEndpoint(e)
	srv.Start()
}
