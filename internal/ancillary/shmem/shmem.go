package shmem

type blockInfo struct {
	data    []byte // allocated shared memory block
	len     uint32
	cap     uint32
	objname string
}

func (b *blockInfo) ObjectName() string {
	return b.objname
}

func (b *blockInfo) Cap() uint32 {
	return b.cap
}

func (b *blockInfo) Len() uint32 {
	return b.len
}

func (b *blockInfo) WithLen(len uint32) {
	b.len = len
}
