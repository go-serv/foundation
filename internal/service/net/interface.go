package net

import (
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/service"
)

type NetworkServiceInterface interface {
	service.BaseServiceInterface
	Service_OnNewSession(req request.RequestInterface) error
	Service_RequestNewSession(req request.RequestInterface) int32
}
