package platform

import (
	"errors"
	"github.com/go-serv/service/pkg/z"
	pf "github.com/go-serv/service/pkg/z/platform"
	"os"
	"sync/atomic"
)

type fs struct {
	owner          z.UniqueId
	spaceAllocated int64
}

func (f *fs) OpenFile(p pf.Pathname, flags int, mode os.FileMode) (fd pf.FileDescriptor, err error) {
	return
}

func (f *fs) CreateZeroFile(p pf.Pathname, size int64, perm uint32) (fd pf.FileDescriptor, err error) {
	var (
		file *os.File
	)
	file, err = os.OpenFile(p.String(), os.O_CREATE|os.O_RDWR, os.FileMode(perm))
	if err != nil {
		return
	}
	err = os.Truncate(p.String(), size)
	if err != nil {
		return
	}
	fd = pf.FileDescriptor{}
	fd.File = file
	atomic.AddInt64(&f.spaceAllocated, size)
	return
}

func (f *fs) CloseFile(pf.FileDescriptor) {

}

func (f *fs) DirectoryExists(path pf.Pathname) bool {
	_, err := os.Stat(path.String())
	return !errors.Is(err, os.ErrNotExist)
}

func (f *fs) CreateDir(path pf.Pathname, perm uint32) (err error) {
	return
}

func (f *fs) RmDir(pf.Pathname) (err error) {
	return
}

func (f *fs) AvailableDiskSpace(p pf.Pathname) uint64 {
	return 0
}
