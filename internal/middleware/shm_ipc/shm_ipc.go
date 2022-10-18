package shm_ipc

//type ipcType struct {
//	memPool *sharedMemPool
//	codec   z.CodecInterface
//}
//
//func (ipc *ipcType) unmarshalHandler(next z.MwChainElement, in []byte, _ z.MethodReflectionInterface, msg z.MessageReflectionInterface, df z.DataFrameInterface) (out []byte, err error) {
//	if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
//		out = in
//		_, err = next(out)
//		if err != nil {
//			return
//		}
//		return
//	}
//	//
//	dfl := df.(z.LocalDataFrameInterface)
//	block := shmem.NewForRead(dfl.SharedMemObjectName(), dfl.SharedMemDataSize(), dfl.SharedMemBlockSize())
//	out, err = block.Read()
//	if err != nil {
//		return
//	}
//	//
//	_, err = next(out)
//	if err != nil {
//		return
//	}
//	return
//}
//
//func (ipc *ipcType) marshalHandler(next z.MwChainElement, in []byte, _ z.MethodReflectionInterface, msg z.MessageReflectionInterface, df z.DataFrameInterface) (out []byte, err error) {
//	if _, has := msg.Get(go_serv.E_LocalShmIpc); !has {
//		out = in
//		_, err = next(out)
//		if err != nil {
//			return
//		}
//		return
//	}
//	//
//	blockSize := uint32(len(in))
//	memBlock := <-ipc.memPool.acquire(blockSize)
//	if memBlock == nil {
//		return nil, fmt.Errorf("shmem pool: failed to acquire a memory block of %d bytes size", blockSize)
//	}
//	//
//	if err = memBlock.Write(in); err != nil {
//		return
//	}
//	dfl := df.(z.LocalDataFrameInterface)
//	dfl.WithSharedMemBlockSize(uint32(len(in)))
//	dfl.WithSharedMemObjectName(memBlock.ObjectName())
//	dfl.WithHeaderFlag(local_cc.SharedMem_IPC)
//	_, err = next(out)
//	if err != nil {
//		return
//	}
//	return
//}
