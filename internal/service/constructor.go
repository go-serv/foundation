package service

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/grpc/descriptor"
)

func newBaseService(name string) baseService {
	s := baseService{name: name}
	s.state = StateInit
	s.sd = descriptor.NewServiceDescriptor(name)
	return s
}

func NewLocalService(name string) *localService {
	s := &localService{newBaseService(name)}
	return s
}

func NewNetworkService(name string) NetworkServiceInterface {
	s := &networkService{newBaseService(name)}
	s.sd.AddMethodProtoExt(go_serv.E_NetNewSession)
	s.sd.Populate()
	return s
}
