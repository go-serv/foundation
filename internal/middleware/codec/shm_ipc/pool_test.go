package shm_ipc

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

type testFixtures struct {
	wg   sync.WaitGroup
	pool *sharedMemPool
}

func (f *testFixtures) worker(t *testing.T) {
	go func() {
		randSize := uint32(1024 + rand.Intn(4096))
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
	for i := 0; i < 900; i++ {
		f.wg.Add(1)
		f.worker(t)
	}
	f.wg.Wait()
}
