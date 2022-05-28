package shmem

import (
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"time"
)

func NewSharedMemory(serviceName protoreflect.FullName, size int) *shmemInfo {
	sh := new(shmemInfo)
	sh.objname = "/dev/shm/" + string(serviceName) + "." + strconv.Itoa(int(time.Now().UnixNano()))
	sh.size = size
	return sh
}

func (a *shmemInfo) Allocate() (err error) {
	var fd int
	fd, err = unix.Open(a.objname, unix.O_CREAT|unix.O_RDWR|unix.O_CLOEXEC, 0600)
	if err != nil {
		return
	}
	// https://man7.org/linux/man-pages/man3/ftruncate.3p.html
	// If fildes refers to a shared memory object, ftruncate() shall set
	// the size of the shared memory object to length.
	err = unix.Ftruncate(fd, int64(a.size))
	if err != nil {
		return
	}
	a.shmem, err = unix.Mmap(fd, 0, a.size, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return
	}
	err = unix.Close(fd)
	return
}

func (a *shmemInfo) Populate(src []byte) error {
	if len(src) > a.size {
		return nil
	}
	n := copy(a.shmem[0:a.size], src)
	if n != a.size {
		return nil
	}
	return nil
}

func (a *shmemInfo) Free() error {
	return unix.Munmap(a.shmem)
}
