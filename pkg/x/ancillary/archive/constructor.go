package archive

import (
	"compress/gzip"
	"github.com/go-serv/service/internal/ancillary/archive"
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/ancillary"
	"github.com/go-serv/service/pkg/z/platform"
)

type ArchiveOptions struct {
	copyable.Shallow
	ancillary.ArchiveOptions
}

func (in ArchiveOptions) NewUntar(target platform.Pathname, comp ancillary.CompressorType) (z.RunnableInterface, error) {
	def := ArchiveOptions{}
	def.GzipMultistream = true
	copyable.Shallow{}.Merge(def, in)
	return archive.NewUntar(target, comp, def.ArchiveOptions)
}

func (in ArchiveOptions) NewTar(target platform.Pathname, comp ancillary.CompressorType) (z.RunnableInterface, error) {
	def := ArchiveOptions{}
	def.GzipLevel = gzip.DefaultCompression
	copyable.Shallow{}.Merge(def, in)
	return archive.NewUntar(target, comp, def.ArchiveOptions)
}
