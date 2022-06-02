package shmem

import (
	"crypto/rand"
	"encoding/binary"
)

type BlockRandId uint64

func (BlockRandId) New() BlockRandId {
	var buf [8]byte
	_, _ = rand.Read(buf[:])
	out := binary.LittleEndian.Uint64(buf[:])
	out = out & (^uint64(0) >> 9)
	return BlockRandId(out)
}

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
