package archive

import (
	platform2 "github.com/go-serv/service/internal/platform"
	"github.com/go-serv/service/pkg/z/platform"
	"io"
)

type archive struct {
	fs           platform.FilesystemInterface
	fsPerms      platform2.UnixPerms
	decompReader io.Reader
	compWriter   io.Writer
	target       platform.Pathname
}
