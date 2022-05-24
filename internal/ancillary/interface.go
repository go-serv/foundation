package ancillary

import "io"

type Net_RW_Info interface {
	Offset() int
}

type NetReaderInterface interface {
	Net_RW_Info
	ReadString() (string, error)
}

type NetWriterInterface interface {
	Net_RW_Info
	io.Writer
	WriteString(string) error
	Bytes() []byte
}
