// build +linux

package shmem

import (
	"crypto/rand"
	"testing"
)

func alloc_failed(t *testing.T, err error) {
	t.Fatalf("failed to allocate shared memory, error: %v", err)
}

func free_failed(t *testing.T, err error) {
	t.Fatalf("failed to free shared memory, error: %v", err)
}

func genRandomData(size int) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func TestAllocateAndFree(t *testing.T) {
	memSize := 1024
	shm := NewSharedMemory("", memSize)
	if err := shm.Allocate(); err != nil {
		alloc_failed(t, err)
	}
	if err := shm.Free(); err != nil {
		free_failed(t, err)
	}
}

func TestRead(t *testing.T) {
	memSize := 1024
	shm := NewSharedMemory("", memSize)
	if err := shm.Allocate(); err != nil {
		alloc_failed(t, err)
	}
	rdata := genRandomData(memSize)
	if err := shm.Write(rdata); err != nil {
		t.Fatalf("failed to populated shared memory with data, error: %v", err)
	}
	// Receiver side
	shmRecv := NewSharedMemory(shm.ObjectName(), shm.Size())
	data, err := shmRecv.Read()
	if err != nil {
		t.Fatalf("failed to read shared memory content, error: %v", err)
	}
	for i := 0; i < memSize; i++ {
		if data[i] != rdata[i] {
			t.Fatalf("memory corrupted")
		}
	}
	if err := shm.Free(); err != nil {
		free_failed(t, err)
	}
}
