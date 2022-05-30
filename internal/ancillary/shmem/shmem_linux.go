package shmem

import (
	"golang.org/x/sys/unix"
	"strconv"
	"time"
)

const UnixPathPrefix = "/dev/shm/go-serv."

var UnixFilePerm uint32 = 0600

func NewSharedMemory(cap uint32) *blockInfo {
	b := new(blockInfo)
	b.objname = UnixPathPrefix + strconv.Itoa(int(time.Now().UnixNano()))
	b.len = cap
	b.cap = cap
	return b
}

func NewForRead(objname string, len uint32, cap uint32) *blockInfo {
	b := new(blockInfo)
	b.objname = objname
	b.len = len
	b.cap = cap
	return b
}

func (b *blockInfo) Allocate() (err error) {
	var fd int
	fd, err = unix.Open(b.objname, unix.O_CREAT|unix.O_RDWR|unix.O_NOFOLLOW, UnixFilePerm)
	if err != nil {
		return
	}
	// https://man7.org/linux/man-pages/man3/ftruncate.3p.html
	// If fildes refers to a shared memory object, ftruncate() shall set
	// the size of the shared memory object to length.
	err = unix.Ftruncate(fd, int64(b.cap))
	if err != nil {
		return
	}
	//
	b.data, err = unix.Mmap(fd, 0, int(b.cap), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return
	}
	err = unix.Close(fd)
	return
}

func (b *blockInfo) Read() (out []byte, err error) {
	var fd int
	fd, err = unix.Open(b.objname, unix.O_RDWR, UnixFilePerm)
	if err != nil {
		return
	}
	b.data, err = unix.Mmap(fd, 0, int(b.len), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return
	}
	_ = unix.Close(fd)
	out = b.data
	return
}

func (b *blockInfo) Write(src []byte) error {
	if len(src) > int(b.cap) {
		return nil
	}
	n := copy(b.data, src)
	if n != len(src) {
		return nil
	}
	return nil
}

func (a *blockInfo) Free() error {
	if err := unix.Munmap(a.data); err != nil {
		return err
	}
	if err := unix.Unlink(a.objname); err != nil {
		return err
	}
	return nil
}
