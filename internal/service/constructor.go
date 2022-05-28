package service

import (
	"fmt"
	"github.com/go-serv/service/internal/grpc/descriptor"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewBaseService(name protoreflect.FullName) *service {
	if name.IsValid() != true {
		panic(fmt.Sprintf("invalid service name '%s'", name))
	}
	s := new(service)
	s.name = name
	s.State = StateInit
	s.sd = descriptor.NewServiceDescriptor(string(name))
	//s.cc = net_cc.NewOrRegistered(name)
	return s
}
