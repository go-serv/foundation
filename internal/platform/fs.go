package platform

import (
	pf "github.com/go-serv/service/pkg/z/platform"
	"os"
)

type fs struct{}

func (f *fs) OpenFile(p pf.Pathname, flags int, mode os.FileMode) (*pf.FileDescriptor, error) {
	return nil, nil
}

func (f *fs) AvailableDiskSpace(p pf.Pathname) uint64 {
	return 0
}
