package server

import (
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/grpc/response"
	"github.com/go-serv/service/internal/service"
)

func NewNetSessionHandler() *mw_group.GroupItem {
	reqHandler := func(r request.RequestInterface, svc interface{}) error {
		netSvc := svc.(service.NetworkServiceInterface)
		return netSvc.Service_OnNewSession(r)
	}
	resHandler := func(r response.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}
