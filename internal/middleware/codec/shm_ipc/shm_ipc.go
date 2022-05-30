package shm_ipc

import (
	"fmt"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary/shmem"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
)

type ipcType struct {
	memPool *sharedMemPool
	codec   i.CodecInterface
}

func (ipc *ipcType) unmarshal(msg i.MessageReflectInterface, df i.LocalDataFrameInterface) (out []byte, err error) {
	block := shmem.NewSharedMemory(df.SharedMemObjectName(), df.SharedMemBlockSize())
	out, err = block.Read()
	if err != nil {
		return nil, err
	}
	err = ipc.codec.PureUnmarshal(out, msg.Value())
	ipc.memPool.release(df.SharedMemObjectName())
	return
}

func (ipc *ipcType) marshal(in []byte, df i.LocalDataFrameInterface) error {
	memBlock := <-ipc.memPool.acquire(len(in))
	if memBlock == nil {
		return fmt.Errorf("failed to")
	}
	//
	if err := memBlock.Write(in); err != nil {
		return err
	}
	df.WithSharedMemBlockSize(len(in))
	df.WithSharedMemObjectName(memBlock.ObjectName())
	return nil
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
