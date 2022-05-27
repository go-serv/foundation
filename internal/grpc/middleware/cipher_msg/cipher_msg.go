package cipher_msg

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/grpc/middleware/mw_group"
)

func NewNetCipherServerHandler() *mw_group.GroupItem {
	reqHandler := func(r i.RequestInterface, svc interface{}) error {
		netSvc := svc.(i.NetworkServiceInterface)
		encrypted := netSvc.Service_InfoMsgEncryption(r.MethodName())
		//un := r.Data().(cc.UnmarshalerInterface)
		// Enforce message encryption
		//if encrypted && !un.DataFrame().HeaderFlags().Has(cc.HeaderFlags32Type(net_cc.Encryption)) {
		//	return status.Error(codes.InvalidArgument, "message must be encrypted")
		//}
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

func NewNetCipherClientHandler(cc i.NetworkClientInterface) *mw_group.GroupItem {
	netSvc := cc.NetService()
	pre := func(next i.MsgProcTaskHandler, in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) (out []byte, err error) {
		out = in
		var encrypted bool
		var name string

		encrypted = netSvc.Service_InfoMsgEncryption(name)
		if encrypted {

		}
		return
	}
	post := func(next i.MsgProcTaskHandler, in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) (out []byte, err error) {
		var name string
		//name, err = msg.MethodName()
		if err != nil {
			return nil, err
		}
		encrypted := netSvc.Service_InfoMsgEncryption(name)
		//if encrypted && !df.HeaderFlags().Has(cc.HeaderFlags32Type(net_cc.Encryption)) {
		//	return nil, status.Error(codes.InvalidArgument, "message must be encrypted")
		//}
		if encrypted {
			out = in
		}
		return
	}
	//
	cc.Codec().ClientProcessor().AddHandlers(pre, post)
	//
	reqHandler := func(r i.RequestInterface, svc interface{}) error {
		return nil
	}
	//
	resHandler := func(r i.ResponseInterface, svc interface{}) error {
		return nil
	}
	return mw_group.NewItem(reqHandler, resHandler)
}
