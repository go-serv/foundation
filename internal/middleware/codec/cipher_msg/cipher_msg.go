package cipher_msg

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary/crypto/aes"
)

// NetServiceInit adds handlers to the default codec middleware group of server
func NetServiceInit(netSvc i.NetworkServiceInterface) {
	unmarshalHandler := func(in []byte, mdReflect i.MethodReflectInterface, msgReflect i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out, _ = enc.Decrypt(in)
		return
	}
	marshalHandler := func(in []byte, mdReflect i.MethodReflectInterface, msgReflect i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out = enc.Encrypt(in)
		return
	}
	//
	netSvc.CodecMiddlewareGroup().AddHandlers(unmarshalHandler, marshalHandler)
}

func NetClientInit(cc i.NetworkClientInterface) {
	netSvc := cc.NetService()
	unmarshalHandler := func(in []byte, mdReflect i.MethodReflectInterface, msgReflect i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out, _ = enc.Decrypt(in)
		return
	}
	marshalHandler := func(in []byte, mdReflect i.MethodReflectInterface, msgReflect i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out = enc.Encrypt(in)
		return
	}
	//
	cc.CodecMiddlewareGroup().AddHandlers(unmarshalHandler, marshalHandler)
}
