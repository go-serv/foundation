package z

import (
	"io"
)

type RunnableInterface interface {
	Run() error
}

type NetReaderInterface interface {
	ReadString() (string, error)
}

type NetWriterInterface interface {
	io.WriteCloser
	WriteString(string) error
	Bytes() []byte
}
