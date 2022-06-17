package platform

import (
	"os"
	"strings"
)

type (
	Pathname string
)

var (
	PathSeparator = string(os.PathSeparator)
)

func (p Pathname) String() string {
	return string(p)
}

func (p Pathname) IsCanonical() bool {
	return true
}

func (p Pathname) ComposePath(parts ...string) Pathname {
	var v string
	path := strings.TrimRight(p.String(), PathSeparator) + PathSeparator
	for i := 0; i < len(parts); i++ {
		if parts[i] == PathSeparator {
			path += PathSeparator
			continue
		}
		v = strings.TrimRight(parts[i], PathSeparator)
		v = strings.TrimLeft(v, PathSeparator)
		if i < len(parts)-1 {
			v += PathSeparator
		}
		path += v
	}
	return Pathname(path)
}

type FileDescriptor struct {
	*os.File
}

type FilesystemInterface interface {
	OpenFile(p Pathname, flags int, mode os.FileMode) (FileDescriptor, error)
	CloseFile(FileDescriptor)
	CreateZeroFile(p Pathname, size int64, perm uint32) (FileDescriptor, error)
	//LockFile(*FileDescriptor) error
	//UnlockFile(*FileDescriptor)
	CreateDir(p Pathname, perm uint32) error
	DirectoryExists(Pathname) bool
	RmDir(Pathname) error
	AvailableDiskSpace(Pathname) uint64
}
