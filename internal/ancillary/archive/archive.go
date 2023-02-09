package archive

import (
	"github.com/mesh-master/foundation/pkg/z/ancillary"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"io"
)

type archive struct {
	ancillary.ArchiveOptions
	fs         platform.FilesystemInterface
	fsPerms    platform.UnixPerms
	compReader io.Reader
	compWriter io.Writer
	target     platform.Pathname
}
