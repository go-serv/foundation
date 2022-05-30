package shmem

type SharedMemoryInterface interface {
	ObjectName() string
	Size() uint32
	Allocate() error
	Read() ([]byte, error)
	Write([]byte) error
	Free() error
}
