package shmem

type SharedMemoryInterface interface {
	ObjectName() string
	Size() int
	Allocate() error
	Read() ([]byte, error)
	Write([]byte) error
	Free() error
}
