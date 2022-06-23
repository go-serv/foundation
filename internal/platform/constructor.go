package platform

import (
	"github.com/go-serv/service/internal/platform/fs"
	"github.com/go-serv/service/pkg/z"
)

func NewPlatform(id z.TenantId) *platform {
	p := new(platform)
	p.FilesystemInterface = fs.NewFilesystem(id)
	return p
}
