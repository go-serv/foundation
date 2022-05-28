package shmem

type shmemInfo struct {
	shmem   []byte // allocated shared memory block
	size    int    // block size
	objname string
}
