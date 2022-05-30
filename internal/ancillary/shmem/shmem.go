package shmem

type blockInfo struct {
	data    []byte // allocated shared memory block
	size    uint32 // block size
	objname string
}

func (b *blockInfo) ObjectName() string {
	return b.objname
}

func (b *blockInfo) Size() uint32 {
	return b.size
}
