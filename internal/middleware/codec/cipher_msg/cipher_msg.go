package cipher_msg

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary/crypto/aes"
)

// NetServiceInit adds handlers to the default codec middleware group of server
func NetServiceInit(netSvc i.NetworkServiceInterface) {
	unmarshalHandler := func(next i.MwChainElement, in []byte, mdReflect i.MethodReflectionInterface, msgReflect i.MessageReflectionInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out, _ = enc.Decrypt(in)
		_, err = next(out)
		return
	}
	marshalHandler := func(next i.MwChainElement, in []byte, mdReflect i.MethodReflectionInterface, msgReflect i.MessageReflectionInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out = enc.Encrypt(in)
		_, err = next(out)
		return
	}
	//
	netSvc.CodecMiddlewareGroup().AddHandlers(unmarshalHandler, marshalHandler)
}

func NetClientInit(cc i.NetworkClientInterface) {
	netSvc := cc.NetService()
	marshalHandler := func(next i.MwChainElement, in []byte, mdReflect i.MethodReflectionInterface, msgReflect i.MessageReflectionInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out = enc.Encrypt(in)
		_, err = next(out)
		return
	}
	unmarshalHandler := func(next i.MwChainElement, in []byte, mdReflect i.MethodReflectionInterface, msgReflect i.MessageReflectionInterface, df i.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out, _ = enc.Decrypt(in)
		_, err = next(out)
		return
	}
	//
	cc.CodecMiddlewareGroup().AddHandlers(unmarshalHandler, marshalHandler)
}
