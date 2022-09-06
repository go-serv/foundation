package archive

import (
	"github.com/go-serv/foundation/pkg/z/ancillary"
	"github.com/go-serv/foundation/pkg/z/platform"
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
