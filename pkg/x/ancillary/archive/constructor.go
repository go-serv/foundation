package archive

import (
	"compress/gzip"
	"github.com/go-serv/service/internal/ancillary/archive"
	"github.com/go-serv/service/internal/ancillary/struc/copyable"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/ancillary"
	"github.com/go-serv/service/pkg/z/platform"
)

type Options struct {
	copyable.Shallow
	ancillary.ArchiveOptions
}

func (in Options) NewUntar(target platform.Pathname, comp ancillary.CompressorType) (z.RunnableInterface, error) {
	def := Options{}
	def.GzipMultistream = true
	def.PlatformOwner = runtime.Runtime().Platform()
	copyable.Shallow{}.Merge(def, in)
	return archive.NewUntar(target, comp, def.ArchiveOptions)
}

func (in Options) NewTar(target platform.Pathname, comp ancillary.CompressorType) (z.RunnableInterface, error) {
	def := Options{}
	def.GzipLevel = gzip.DefaultCompression
	copyable.Shallow{}.Merge(def, in)
	return archive.NewUntar(target, comp, def.ArchiveOptions)
}
