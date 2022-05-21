package ancillary

import "io"

type NetReaderInterface interface {
	ReadString() (string, error)
}

type NetWriterInterface interface {
	io.Writer
	WriteString(string) error
	Bytes() []byte
}
