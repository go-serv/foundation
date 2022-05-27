package service

import (
	"github.com/go-serv/service/internal/grpc/descriptor"
)

func NewBaseService(name string) BaseService {
	s := BaseService{Name: name}
	s.State = StateInit
	s.Sd = descriptor.NewServiceDescriptor(name)
	//s.cc = net_cc.NewOrRegistered(name)
	return s
}
