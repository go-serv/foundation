package net

import (
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/server"
)

type netServer struct {
	server.Server
	mwGroup mw_group.MiddlewareGroupInterface
}

func (s *netServer) sessionMwHandler(r request.RequestInterface) error {
	return nil
}
