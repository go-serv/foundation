package fs

import (
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/platform"
)

func NewFilesystem(tenantId z.TenantId) *filesystem {
	fs := new(filesystem)
	fs.tenantId = tenantId
	return fs
}

func NewWalker(dirHandler platform.DirTreeWalkStepFn, regFileHandler platform.DirTreeWalkStepFn) recursiveWalk {
	rw := recursiveWalk{}
	rw.dirHandler = dirHandler
	rw.regFileHandler = regFileHandler
	return rw
}
