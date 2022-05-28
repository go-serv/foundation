package shmem

type SharedMemoryInterface interface {
	Name() string
	Size() int
	Allocate() error
	Populate([]byte) error
	Read() ([]byte, error)
	Free() error
}
