package shm_ipc

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

const (
	minBlockSize = 1 * 1024 * 1024
	maxBlockSize = 4 * 1024 * 1024
)

type testFixtures struct {
	wg   sync.WaitGroup
	pool *sharedMemPool
}

func (f *testFixtures) worker(t *testing.T) {
	go func() {
		randSize := uint32(minBlockSize + rand.Intn(maxBlockSize-minBlockSize))
		block := <-f.pool.acquire(randSize)
		if block == nil {
			t.Fatalf("failed to acquire memory block")
		}
		// do something with the acquired memory block
		time.Sleep(5 * time.Millisecond)
		f.pool.release(block.ObjectName())
		f.wg.Done()
	}()
}

func TestMemPool(t *testing.T) {
	f := new(testFixtures)
	f.wg = sync.WaitGroup{}
	f.pool = newSharedMemPool(5)
	for i := 0; i < 100; i++ {
		f.wg.Add(1)
		f.worker(t)
	}
	f.wg.Wait()
}
