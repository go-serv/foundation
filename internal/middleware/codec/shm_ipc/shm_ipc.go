package shm_ipc

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
)

func Init(cc i.LocalClientInterface) {
	unmarshalHandler := func(in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := md.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	marshalHandler := func(in []byte, md i.MethodDescriptorInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := md.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	//
	mg := cc.CodecMiddlewareGroup()
	mg.AddHandlers(unmarshalHandler, marshalHandler)
}
