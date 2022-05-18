package pkg

import (
	"github.com/go-serv/service/internal/autogen/proto/net"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type serviceInterface interface {
	Name() string
	AddProtoExtension(info *protoimpl.ExtensionInfo, defaultValue interface{})
}

type NetworkServiceApi interface {
	serviceInterface
	net.NetParcelServer
}

type LocalServiceApi interface {
	serviceInterface
}
