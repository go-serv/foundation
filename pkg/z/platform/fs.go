package platform

import "os"

type (
	Pathname string
)

type FileDescriptor struct {
	*os.File
}

type FilesystemInterface interface {
	OpenFile(p Pathname, flags int, mode os.FileMode) (*FileDescriptor, error)
	//CloseFile(*FileDescriptor)
	//LockFile(*FileDescriptor) error
	//UnlockFile(*FileDescriptor)
	//CreateDir(Pathname) error
	//RmDir(Pathname) error
	AvailableDiskSpace(Pathname) uint64
}
