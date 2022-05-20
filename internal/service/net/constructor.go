package net

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/service"
)

func NewNetworkService(name string) *networkService {
	s := &networkService{service.NewBaseService(name)}
	s.Sd.AddServiceProtoExt(go_serv.E_NetMsgEnc)
	s.Sd.AddMethodProtoExt(go_serv.E_NetNewSession)
	s.Sd.AddMethodProtoExt(go_serv.E_MNetMsgEnc)
	s.Sd.Populate()
	return s
}
