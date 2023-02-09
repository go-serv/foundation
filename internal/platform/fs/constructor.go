package fs

import (
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/platform"
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
