package shm_ipc

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	local_cc "github.com/go-serv/service/internal/grpc/codec/local"
)

func ClientInit(cc i.LocalClientInterface) {
	ipc := new(ipcType)
	ipc.memPool = newSharedMemPool(50)
	ipc.codec = cc.Codec()
	// Handlers
	unmarshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		out, err = ipc.unmarshal(in, msg, df.(i.LocalDataFrameInterface))
		return
	}
	marshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		df.WithHeaderFlag(local_cc.SharedMemRead)
		out, err = ipc.marshal(in, df.(i.LocalDataFrameInterface))
		return
	}
	//
	mg := cc.CodecMiddlewareGroup()
	mg.AddHandlers(unmarshalHandler, marshalHandler)
}

func ServiceInit(srv i.LocalServiceInterface) {
	ipc := new(ipcType)
	ipc.memPool = newSharedMemPool(5)
	ipc.codec = srv.Codec()
	//
	unmarshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		out, err = ipc.unmarshal(in, msg, df.(i.LocalDataFrameInterface))
		return
	}
	marshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		out, err = ipc.marshal(in, df.(i.LocalDataFrameInterface))
		return
	}
	//
	mg := srv.CodecMiddlewareGroup()
	mg.AddHandlers(unmarshalHandler, marshalHandler)
}
