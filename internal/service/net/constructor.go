package net

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	"github.com/go-serv/service/internal/middleware/codec/cipher_msg"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewNetworkService(name protoreflect.FullName) *netService {
	s := &netService{}
	s.ServiceInterface = service.NewBaseService(name)
	cc := net_cc.NewOrRegistered(string(name))
	s.ServiceInterface.WithCodec(cc)
	s.ServiceInterface.Service_AddServiceProtoExtension(go_serv.E_NetMsgEnc)
	//s.sd.AddServiceProtoExt(go_serv.E_NetMsgEnc)
	//s.sd.AddMethodProtoExt(go_serv.E_NetNewSession)
	//s.sd.AddMethodProtoExt(go_serv.E_MNetMsgEnc)
	s.ServiceInterface.Service_Descriptor().Populate()
	cipher_msg.NetServiceInit(s)
	return s
}
