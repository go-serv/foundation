package shm_ipc

import (
	"fmt"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary/shmem"
)

type ipcType struct {
	memPool *sharedMemPool
	codec   i.CodecInterface
}

func (ipc *ipcType) unmarshal(in []byte, msg i.MessageReflectInterface, df i.LocalDataFrameInterface) (out []byte, err error) {
	block := shmem.NewSharedMemory(df.SharedMemObjectName(), df.SharedMemBlockSize())
	out, err = block.Read()
	if err != nil {
		return nil, err
	}
	err = ipc.codec.PureUnmarshal(out, msg.Value())
	ipc.memPool.release(df.SharedMemObjectName())
	return
}

func (ipc *ipcType) marshal(in []byte, df i.LocalDataFrameInterface) (out []byte, err error) {
	memBlock := <-ipc.memPool.acquire(uint32(len(in)))
	if memBlock == nil {
		return nil, fmt.Errorf("failed to")
	}
	//
	if err = memBlock.Write(in); err != nil {
		return
	}
	df.WithSharedMemBlockSize(uint32(len(in)))
	df.WithSharedMemObjectName(memBlock.ObjectName())
	out, err = df.Compose(nil)
	return
}
