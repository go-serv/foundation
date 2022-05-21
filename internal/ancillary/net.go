package ancillary

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
)

type net_IO_Types interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | uintptr | bool | float32 | float64
}

type Net_IO interface {
}

type netReader struct {
	buf []byte
}

type netWriter struct {
	buf *bytes.Buffer
}

func NewNetReader(data []byte) *netReader {
	r := new(netReader)
	r.buf = data
	return r
}

func NewNetWriter() *netWriter {
	w := new(netWriter)
	w.buf = new(bytes.Buffer)
	return w
}

func (w *netWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func NetReader[T net_IO_Types](r *netReader) (T, error) {
	var out T
	n := reflect.TypeOf(out).Size()
	if len(r.buf) < int(n) {
		return out, io.EOF
	}
	err := binary.Read(bytes.NewReader(r.buf[:n]), binary.BigEndian, &out)
	if err != nil {
		return out, err
	}
	r.buf = r.buf[n:]
	return out, nil
}

func NetWriter[T net_IO_Types](w *netWriter, t T) error {
	err := binary.Write(w.buf, binary.BigEndian, t)
	if err != nil {
		return err
	} else {
		return nil
	}
}
