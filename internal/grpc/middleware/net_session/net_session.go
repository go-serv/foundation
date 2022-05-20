package net_session

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
)

func NewNetSessionHandler() *mw_group.GroupItem {
	reqHandler := func(r i.RequestInterface, svc interface{}) error {
		netSvc := svc.(i.NetworkServiceInterface)
		timeout := netSvc.Service_InfoNewSession(r.MethodName())
		if timeout != 0 { // Create new session
			netSvc.Service_OnNewSession(r)
			return nil
		}
		return nil
	}
	resHandler := func(r i.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}
