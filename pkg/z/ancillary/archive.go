package ancillary

import "github.com/go-serv/service/pkg/z/platform"

type CompressorType int

const (
	GzipCompressor CompressorType = iota
)

type ArchiveOptions struct {
	PlatformOwner   platform.FilesystemInterface
	GzipMultistream bool
	GzipLevel       int
}
