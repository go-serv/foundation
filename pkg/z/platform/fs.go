package platform

import (
	"os"
)

type (
	Pathname          string
	UnixPerms         uint32
	DirTreeWalkStepFn func(path string, info os.FileInfo) error
)

var (
	PathSeparator = string(os.PathSeparator)
)

type FilesystemInterface interface {
	AvailableDiskSpace(Pathname) uint64
	Write(fd *os.File, data []byte) error
	WriteAt(fd *os.File, off int64, data []byte) error
	CreateZeroFile(p Pathname, size int64, perm UnixPerms) (*os.File, error)
	CreateDir(p Pathname, perms UnixPerms) error
	Remove(Pathname) error
}
