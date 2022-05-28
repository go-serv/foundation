// build +linux

package shmem

import (
	"crypto/rand"
	"testing"
)

func genRandomData(size int) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func TestAllocate(t *testing.T) {
	memSize := 1024
	shm := NewSharedMemory("myservice", memSize)
	if err := shm.Allocate(); err != nil {
		t.Fatalf("failed to allocate shared memory, error: %v", err)
	}
	rdata := genRandomData(memSize)
	if err := shm.Populate(rdata); err != nil {
		t.Fatalf("failed to populated shared memory with data, error: %v", err)
	}
	if err := shm.Free(); err != nil {
		t.Fatalf("failed to free shared memory, error: %v", err)
	}
}
