package net

import (
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewNetworkService(name protoreflect.FullName) *netService {
	s := &netService{}
	s.ServiceInterface = service.NewBaseService(name)
	cc := net_cc.NewOrRegistered(string(name))
	s.ServiceInterface.WithCodec(cc)
	return s
}
