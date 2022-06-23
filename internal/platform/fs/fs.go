package fs

import (
	"github.com/go-serv/service/pkg/z"
	pf "github.com/go-serv/service/pkg/z/platform"
	"os"
	"path/filepath"
	"sync/atomic"
)

type filesystem struct {
	tenantId  z.TenantId
	spaceUsed int64
}

type recursiveWalk struct {
	target         pf.Pathname
	dirHandler     pf.DirTreeWalkStepFn
	regFileHandler pf.DirTreeWalkStepFn
}

func (rw recursiveWalk) walkFn(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	switch m := info.Mode(); true {
	case m.IsDir():
		if rw.dirHandler != nil {
			if err = rw.dirHandler(path, info); err != nil {
				return err
			}
		}
		if err = filepath.Walk(path, rw.walkFn); err != nil {
			return err
		}
	case m.IsRegular():
		if rw.regFileHandler != nil {
			if err = rw.regFileHandler(path, info); err != nil {
				return err
			}
		}
	}
	return nil
}

func (rw recursiveWalk) Run() error {
	return filepath.Walk(rw.target.String(), rw.walkFn)
}

func (f *filesystem) Write(fd *os.File, data []byte) (err error) {
	var n int
	if n, err = fd.Write(data); err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, int64(n))
	return
}

func (f *filesystem) WriteAt(fd *os.File, offset int64, data []byte) (err error) {
	var n int
	if n, err = fd.WriteAt(data, offset); err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, int64(n))
	return
}

func (f *filesystem) CreateZeroFile(p pf.Pathname, size int64, perms pf.UnixPerms) (fd *os.File, err error) {
	if fd, err = os.OpenFile(p.String(), os.O_CREATE|os.O_RDWR, os.FileMode(perms)); err != nil {
		return
	}
	if err = os.Truncate(p.String(), size); err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, size)
	return
}

func (f *filesystem) Remove(target pf.Pathname) (err error) {
	var (
		unallocated int64
		info        os.FileInfo
	)
	if info, err = os.Stat(target.String()); err != nil {
		return
	}
	// Calculate the total size of disk space that will be fred if target to be removed is a directory.
	switch m := info.Mode(); true {
	case m.IsDir():
		walker := NewWalker(nil, func(path string, info os.FileInfo) error {
			unallocated += info.Size()
			return nil
		})
		err = walker.Run()
	case m.IsRegular():
		unallocated = info.Size()
	}
	if err = os.Remove(target.String()); err != nil {
		return
	}
	atomic.AddInt64(&f.spaceUsed, -unallocated)
	return
}

func (f *filesystem) CreateDir(path pf.Pathname, perms pf.UnixPerms) (err error) {
	err = os.MkdirAll(path.String(), os.FileMode(perms))
	return
}

func (f *filesystem) RemoveDir(pf.Pathname) (err error) {
	return
}

func (f *filesystem) AvailableDiskSpace(p pf.Pathname) uint64 {
	return 0
}
