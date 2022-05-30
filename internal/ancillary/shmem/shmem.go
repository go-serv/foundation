package shmem

type blockInfo struct {
	data    []byte // allocated shared memory block
	size    int    // block size
	objname string
}

func (b *blockInfo) ObjectName() string {
	return b.objname
}

func (b *blockInfo) Size() int {
	return b.size
}
