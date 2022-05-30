package shm_ipc

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary/shmem"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
)

type ipcType struct {
	memPool *sharedMemPool
	codec   i.CodecInterface
}

func (ipc *ipcType) unmarshal(msg i.MessageReflectInterface, df i.LocalDataFrameInterface) (err error) {
	var ipcData []byte
	block := shmem.NewSharedMemory(df.SharedMemObjectName(), df.SharedMemBlockSize())
	ipcData, err = block.Read()
	if err != nil {
		return err
	}
	return ipc.codec.PureUnmarshal(ipcData, msg.Value())
}

func ServiceInit(srv i.LocalServiceInterface) {
	unmarshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	marshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
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
	unmarshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	marshalHandler := func(in []byte, method i.MethodReflectInterface, msg i.MessageReflectInterface, df i.DataFrameInterface) (out []byte, err error) {
		if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
			out = in
			return
		}
		return
	}
	//
	mg := cc.CodecMiddlewareGroup()
	mg.AddHandlers(unmarshalHandler, marshalHandler)
}
