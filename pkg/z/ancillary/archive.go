package ancillary

type CompressorType int

const (
	GzipCompressor CompressorType = iota
)

type ArchiveOptions struct {
	GzipMultistream bool
	GzipLevel       int
}
