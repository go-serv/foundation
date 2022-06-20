package platform

import (
	"errors"
	"github.com/go-serv/service/pkg/z"
	pf "github.com/go-serv/service/pkg/z/platform"
	"os"
	"sync/atomic"
)

type Filesystem struct {
	owner     z.UniqueId
	spaceUsed int64
}

func (f *Filesystem) OpenFile(p pf.Pathname, flags int, mode os.FileMode) (fd pf.FileDescriptor, err error) {
	return
}

func (f *Filesystem) Write(fd *os.File, data []byte) (err error) {
	var n int
	if n, err = fd.Write(data); err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, int64(n))
	return
}

func (f *Filesystem) WriteAt(fd *os.File, offset int64, data []byte) (err error) {
	var n int
	if n, err = fd.WriteAt(data, offset); err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, int64(n))
	return
}

func (f *Filesystem) CreateZeroFile(p pf.Pathname, size int64, perms pf.UnixPerms) (fd *os.File, err error) {
	fd, err = os.OpenFile(p.String(), os.O_CREATE|os.O_RDWR, os.FileMode(perms))
	if err != nil {
		return
	}
	err = os.Truncate(p.String(), size)
	if err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, int64(size))
	return
}

func (f *Filesystem) CloseFile(pf.FileDescriptor) {

}

func (f *Filesystem) DirectoryExists(path pf.Pathname) bool {
	_, err := os.Stat(path.String())
	return !errors.Is(err, os.ErrNotExist)
}

func (f *Filesystem) CreateDir(path pf.Pathname, perms pf.UnixPerms) (err error) {
	err = os.MkdirAll(path.String(), os.FileMode(perms))
	return
}

func (f *Filesystem) RmDir(pf.Pathname) (err error) {
	return
}

func (f *Filesystem) AvailableDiskSpace(p pf.Pathname) uint64 {
	return 0
}
