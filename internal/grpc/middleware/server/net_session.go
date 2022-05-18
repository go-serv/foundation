package server

import (
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"github.com/go-serv/service/internal/grpc/request"
)

func NewNetSessionHandler() mw_group.MiddlewareHandler {
	return func(r request.RequestInterface) error {
		return nil
	}
}
