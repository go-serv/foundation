package shmem

type SharedMemoryInterface interface {
	ObjectName() string
	Cap() uint32
	Len() uint32
	Allocate() error
	Read() ([]byte, error)
	Write([]byte) error
	Free() error
}
