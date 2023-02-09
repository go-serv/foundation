package platform

import (
	"github.com/mesh-master/foundation/internal/platform/fs"
	"github.com/mesh-master/foundation/pkg/z"
)

func NewPlatform(id z.TenantId) *platform {
	p := new(platform)
	p.FilesystemInterface = fs.NewFilesystem(id)
	return p
}
