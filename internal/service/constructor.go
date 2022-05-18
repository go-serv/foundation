package service

import "github.com/go-serv/service/internal/autogen/proto/go_serv"

func newBaseService(name string) baseService {
	s := baseService{name: name}
	s.state = StateInit
	s.Service_AddMethodProtoExtension(go_serv.E_NewSession, false)
	return s
}

func NewLocalService(name string) *localService {
	s := &localService{newBaseService(name)}
	return s
}

func NewNetworkService(name string) NetworkServiceInterface {
	s := &networkService{newBaseService(name)}
	return s
}
