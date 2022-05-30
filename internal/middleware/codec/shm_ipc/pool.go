package shm_ipc

import (
	"github.com/go-serv/service/internal/ancillary/shmem"
	"runtime"
	"sync"
	"sync/atomic"
)

type memChan chan shmem.SharedMemoryInterface

type memBlockInfo struct {
	block shmem.SharedMemoryInterface
	free  bool
}

type sharedMemPool struct {
	mu     sync.Mutex
	blocks []*memBlockInfo
	inUse  int32
	max    int32
}

func newSharedMemPool(size uint) *sharedMemPool {
	p := new(sharedMemPool)
	p.max = int32(size)
	p.blocks = make([]*memBlockInfo, 0, size)
	return p
}

func (pool *sharedMemPool) release(objname string) {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	for _, v := range pool.blocks {
		if v.block.ObjectName() == objname {
			v.free = true
			atomic.AddInt32(&pool.inUse, -1)
			return
		}
	}
}

func (pool *sharedMemPool) acquire(size uint32) memChan {
	ch := make(memChan, 0)
	go func() {
		if atomic.LoadInt32(&pool.inUse) < pool.max {
			atomic.AddInt32(&pool.inUse, 1)
			pool.mu.Lock()
			defer pool.mu.Unlock()
			mblock := shmem.NewSharedMemory(size)
			err := mblock.Allocate()
			if err != nil {
				atomic.AddInt32(&pool.inUse, -1)
				close(ch)
			} else {
				info := &memBlockInfo{}
				info.block = mblock
				pool.blocks = append(pool.blocks, info)
				ch <- mblock
			}
			return
		} else { // try to acquire any free block in loop
			for {
				pool.mu.Lock()
				smallestFree := -1
				for ii, b := range pool.blocks {
					if !b.free {
						continue
					}
					if b.block.Cap() >= size {
						b.free = false
						b.block.WithLen(size)
						atomic.AddInt32(&pool.inUse, 1)
						ch <- b.block
						goto exit
					} else {
						if smallestFree == -1 {
							smallestFree = ii
						} else if b.block.Cap() < pool.blocks[smallestFree].block.Cap() {
							smallestFree = ii
						}
					}
				}
				// Evict the smallest one free block
				if smallestFree != -1 {
					if err := pool.blocks[smallestFree].block.Free(); err != nil {
						// Remove from the poll
						atomic.AddInt32(&pool.inUse, -1)
						pool.blocks = append(pool.blocks[:smallestFree], pool.blocks[smallestFree+1:]...)
						close(ch)
						goto exit
					}
					mblock := shmem.NewSharedMemory(size)
					err := mblock.Allocate()
					if err != nil {
						// The block was freed but failed to be re-allocated
						// Remove from the pool
						atomic.AddInt32(&pool.inUse, -1)
						pool.blocks = append(pool.blocks[:smallestFree], pool.blocks[smallestFree+1:]...)
						close(ch)
						goto exit
					}
					pool.blocks[smallestFree].free = false
					pool.blocks[smallestFree].block = mblock
					ch <- mblock
					goto exit
				}
				pool.mu.Unlock()
				// Wait until any of the blocks will be released
				if atomic.LoadInt32(&pool.inUse) == pool.max {
					runtime.Gosched()
				}
			}
		exit:
			pool.mu.Unlock()
		}
	}()
	return ch
}
