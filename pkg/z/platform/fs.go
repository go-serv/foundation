package platform

import (
	"os"
	"strings"
)

type (
	Pathname  string
	UnixPerms uint32
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

func (p Pathname) FileExists() bool {
	_, err := os.Stat(p.String())
	if err != nil {
		return false
	} else {
		return true
	}
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
	Write(fd *os.File, data []byte) error
	WriteAt(fd *os.File, off int64, data []byte) error
	CloseFile(FileDescriptor)
	CreateZeroFile(p Pathname, size int64, perm UnixPerms) (*os.File, error)
	//LockFile(*FileDescriptor) error
	//UnlockFile(*FileDescriptor)
	CreateDir(p Pathname, perms UnixPerms) error
	DirectoryExists(Pathname) bool
	RmDir(Pathname) error
	AvailableDiskSpace(Pathname) uint64
}
