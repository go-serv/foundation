package shm_ipc

import (
	i "github.com/go-serv/service/internal"
)

func ClientInit(cc i.LocalClientInterface) {
	ipc := new(ipcType)
	ipc.memPool = newSharedMemPool(50)
	ipc.codec = cc.Codec()
	mg := cc.CodecMiddlewareGroup()
	mg.AddHandlers(ipc.unmarshalHandler, ipc.marshalHandler)
}

func ServiceInit(srv i.LocalServiceInterface) {
	ipc := new(ipcType)
	ipc.memPool = newSharedMemPool(5)
	ipc.codec = srv.Codec()
	mg := srv.CodecMiddlewareGroup()
	mg.AddHandlers(ipc.unmarshalHandler, ipc.marshalHandler)
}
