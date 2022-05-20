package net

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/service"
)

func NewNetworkService(name string) NetworkServiceInterface {
	s := &networkService{service.NewBaseService(name)}
	s.Sd.AddMethodProtoExt(go_serv.E_NetNewSession)
	s.Sd.Populate()
	return s
}
