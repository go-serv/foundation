package shmem

type SharedMemoryInterface interface {
	Allocate() error
	Populate([]byte) error
	Free() error
}
