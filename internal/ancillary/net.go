package ancillary

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
)

type net_IO_Types interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | uintptr | bool | float32 | float64
}

type netReader struct {
	buf []byte
}

// reads a length-prefixed string
func (r *netReader) ReadString() (string, error) {
	var sl uint32
	var err error
	sl, err = NetReader[uint32](r)
	if err != nil {
		return "", err
	}
	if sl > uint32(len(r.buf)) {
		return "", Error_IO_OutOfRange
	} else {
		s := make([]byte, sl)
		copy(s, r.buf[0:sl])
		r.buf = r.buf[sl:]
		return string(s), nil
	}
}

type netWriter struct {
	buf *bytes.Buffer
}

func (w *netWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func (w *netWriter) Write(data []byte) (int, error) {
	return w.buf.Write(data)
}

func (w *netWriter) WriteString(s string) error {
	if err := NetWriter[uint32](w, uint32(len(s))); err != nil {
		return err
	}
	_, err := w.Write([]byte(s))
	return err
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
