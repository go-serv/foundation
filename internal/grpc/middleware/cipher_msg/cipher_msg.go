package cipher_msg

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"google.golang.org/protobuf/proto"
)

type encryptedMessage struct {
	proto.Message
	nonce []byte
	key   []byte
}

func NewNetCipherServerHandler() *mw_group.GroupItem {
	reqHandler := func(r i.RequestInterface, svc interface{}) error {
		netSvc := svc.(i.NetworkServiceInterface)
		encrypted := netSvc.Service_InfoMsgEncryption(r.MethodName())
		if encrypted {
			//data := r.Data()
			//msg := &encryptedMessage{}
			//msg.Message = data.(proto.Message)
			//msg.nonce = []byte{1}
			//msg.key = []byte{1}
			//r.WithData(msg)
		}
		return nil
	}
	resHandler := func(r i.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}

func NewNetCipherClientHandler() *mw_group.GroupItem {
	reqHandler := func(r i.RequestInterface, svc interface{}) error {
		netSvc := svc.(i.NetworkServiceInterface)
		encrypted := netSvc.Service_InfoMsgEncryption(r.MethodName())
		if encrypted {
			msg := &encryptedMessage{}
			msg.Message = r.Data().(proto.Message)
			msg.nonce = []byte{1}
			msg.key = []byte{1}
			r.WithData(msg)
		}
		return nil
	}
	resHandler := func(r i.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}
