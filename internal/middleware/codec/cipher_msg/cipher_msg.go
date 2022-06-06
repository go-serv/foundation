package cipher_msg

import (
	"github.com/go-serv/service/internal/ancillary/crypto/aes"
	"github.com/go-serv/service/pkg/z"
)

// NetServiceInit adds handlers to the default codec middleware group of server
func NetServiceInit(netSvc z.NetworkServiceInterface) {
	unmarshalHandler := func(next z.MwChainElement, in []byte, mdReflect z.MethodReflectionInterface, msgReflect z.MessageReflectionInterface, df z.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out, _ = enc.Decrypt(in)
		_, err = next(out)
		return
	}
	marshalHandler := func(next z.MwChainElement, in []byte, mdReflect z.MethodReflectionInterface, msgReflect z.MessageReflectionInterface, df z.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out = enc.Encrypt(in)
		_, err = next(out)
		return
	}
	//
	netSvc.CodecMiddlewareGroup().AddHandlers(unmarshalHandler, marshalHandler)
}

func NetClientInit(cc z.NetworkClientInterface) {
	netSvc := cc.NetService()
	marshalHandler := func(next z.MwChainElement, in []byte, mdReflect z.MethodReflectionInterface, msgReflect z.MessageReflectionInterface, df z.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out = enc.Encrypt(in)
		_, err = next(out)
		return
	}
	unmarshalHandler := func(next z.MwChainElement, in []byte, mdReflect z.MethodReflectionInterface, msgReflect z.MessageReflectionInterface, df z.DataFrameInterface) (out []byte, err error) {
		enc, _ := aes.NewCipher(netSvc.EncriptionKey(), []byte{0x1, 0x2, 0x3, 0x4})
		out, _ = enc.Decrypt(in)
		_, err = next(out)
		return
	}
	//
	cc.CodecMiddlewareGroup().AddHandlers(unmarshalHandler, marshalHandler)
}
