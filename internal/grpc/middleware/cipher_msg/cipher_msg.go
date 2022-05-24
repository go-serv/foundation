package cipher_msg

import (
	i "github.com/go-serv/service/internal"
	cc "github.com/go-serv/service/internal/grpc/codec"
	net_cc "github.com/go-serv/service/internal/grpc/codec/net"
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		un := r.Data().(cc.UnmarshalerInterface)
		// Enforce message encryption
		if encrypted && !un.DataFrame().HeaderFlags().Has(cc.HeaderFlags32Type(net_cc.Encryption)) {
			return status.Error(codes.InvalidArgument, "message must be encrypted")
		}
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
			marshaler := r.Data().(cc.MarshalerInterface)
			marshaler.ChainInterceptorHandler(func(data []byte) ([]byte, error) {
				key := netSvc.Service_EncriptionKey()
				nonce := []byte{1}
				_, _ = key, nonce
				marshaler.DataFrame().WithHeaderFlag(cc.HeaderFlags32Type(net_cc.Encryption))
				return data, nil
			})
		}
		return nil
	}
	resHandler := func(r i.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}
