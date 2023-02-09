package archive

import (
	"compress/gzip"
	"github.com/mesh-master/foundation/internal/ancillary/archive"
	"github.com/mesh-master/foundation/pkg/ancillary/struc/copyable"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/ancillary"
	"github.com/mesh-master/foundation/pkg/z/platform"
)

type Options struct {
	copyable.Shallow
	ancillary.ArchiveOptions
	platform.FilesystemInterface
}

func (in Options) NewUntar(target platform.Pathname, comp ancillary.CompressorType) (z.RunnableInterface, error) {
	def := Options{}
	def.GzipMultistream = true
	//def.FilesystemInterface = runtime.Runtime().Platform()
	copyable.Shallow{}.Merge(def, in)
	return archive.NewUntar(def.FilesystemInterface, target, comp, def.ArchiveOptions)
}

func (in Options) NewTar(target platform.Pathname, comp ancillary.CompressorType) (z.RunnableInterface, error) {
	def := Options{}
	def.GzipLevel = gzip.DefaultCompression
	//def.FilesystemInterface = runtime.Runtime().Platform()
	copyable.Shallow{}.Merge(def, in)
	return archive.NewUntar(def.FilesystemInterface, target, comp, def.ArchiveOptions)
}
