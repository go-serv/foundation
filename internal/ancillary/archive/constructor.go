package archive

import (
	"compress/gzip"
	"github.com/go-serv/service/pkg/z/ancillary"
	"github.com/go-serv/service/pkg/z/platform"
	"os"
)

func NewTar(target platform.Pathname, comp ancillary.CompressorType, opts ancillary.ArchiveOptions) (*ttar, error) {
	var (
		fd  *os.File
		err error
	)
	un := new(ttar)
	un.target = target
	un.ArchiveOptions = opts
	if fd, err = os.OpenFile(target.String(), os.O_RDWR, os.FileMode(0555)); err != nil {
		return nil, err
	}
	switch comp {
	case ancillary.GzipCompressor:
		if un.compWriter, err = gzip.NewWriterLevel(fd, un.GzipLevel); err != nil {
			return nil, err
		}
	}
	return un, nil
}

func NewUntar(target platform.Pathname, comp ancillary.CompressorType, opts ancillary.ArchiveOptions) (*ttar, error) {
	var (
		fd  *os.File
		err error
	)
	un := new(ttar)
	un.target = target
	un.ArchiveOptions = opts
	if fd, err = os.OpenFile(target.String(), os.O_RDONLY, os.FileMode(0444)); err != nil {
		return nil, err
	}
	switch comp {
	case ancillary.GzipCompressor:
		if un.compReader, err = gzip.NewReader(fd); err != nil {
			return nil, err
		}
		un.compReader.(*gzip.Reader).Multistream(un.GzipMultistream)
	}
	return un, nil
}
