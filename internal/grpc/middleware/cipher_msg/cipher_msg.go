package cipher_msg

import (
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/grpc/response"
	"github.com/go-serv/service/internal/service/net"
	"google.golang.org/protobuf/proto"
)

type encryptedMessage struct {
	proto.Message
	nonce []byte
	key   []byte
}

func NewNetCipherServerHandler() *mw_group.GroupItem {
	reqHandler := func(r request.RequestInterface, svc interface{}) error {
		netSvc := svc.(net.NetworkServiceInterface)
		encrypted := netSvc.Service_InfoMsgEncryption(r.MethodName())
		if encrypted {
			data := r.Data()
			msg := &encryptedMessage{}
			msg.Message = data.(proto.Message)
			msg.nonce = []byte{1}
			msg.key = []byte{1}
			r.WithData(msg)
		}
		return nil
	}
	resHandler := func(r response.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}

func NewNetCipherClientHandler() *mw_group.GroupItem {
	reqHandler := func(r request.RequestInterface, svc interface{}) error {
		netSvc := svc.(net.NetworkServiceInterface)
		encrypted := netSvc.Service_InfoMsgEncryption(r.MethodName())
		if encrypted {
			msg := 1
			_ = msg
		}
		return nil
	}
	resHandler := func(r response.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}
