package main

import (
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/server"
	"github.com/go-serv/service/internal/server/local"
	_ "github.com/go-serv/service/testsuite/local/server/internal"
)

func main() {
	svc := runtime.Runtime().LocalService()
	e := server.NewLocalEndpoint(svc)
	srv := local.NewServer()
	srv.AddEndpoint(e)
	srv.Start()
}
