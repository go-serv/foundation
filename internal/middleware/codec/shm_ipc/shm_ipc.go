package shm_ipc

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
)

func ServiceInit(srv i.LocalServiceInterface) {
	unmarshalHandler := func(in []byte, mr i.MethodReflectInterface, msgr i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msgr.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	marshalHandler := func(in []byte, mr i.MethodReflectInterface, msgr i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msgr.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	//
	mg := srv.CodecMiddlewareGroup()
	mg.AddHandlers(unmarshalHandler, marshalHandler)
}

func ClientInit(cc i.LocalClientInterface) {
	unmarshalHandler := func(in []byte, mr i.MethodReflectInterface, msgr i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msgr.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	marshalHandler := func(in []byte, mr i.MethodReflectInterface, msgr i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msgr.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	//
	mg := cc.CodecMiddlewareGroup()
	mg.AddHandlers(unmarshalHandler, marshalHandler)
}
