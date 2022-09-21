package io

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"reflect"
)

var (
	ErrOutOfRange = errors.New("Network IO, out of range")
)

type NetRW_Typ interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | uintptr | bool | float32 | float64
}

type netRW_Info struct {
	offset int
}

func (o netRW_Info) Offset() int {
	return o.offset
}

type NetReader struct {
	netRW_Info
	buf []byte
}

// ReadString reads a length-prefixed string
func (r *NetReader) ReadString() (string, error) {
	var sl uint32
	var err error
	sl, err = GenericNetReader[uint32](r)
	if err != nil {
		return "", err
	}
	if sl > uint32(len(r.buf)) {
		return "", ErrOutOfRange
	} else {
		s := make([]byte, sl)
		r.offset += int(sl)
		copy(s, r.buf[0:sl])
		r.buf = r.buf[sl:]
		return string(s), nil
	}
}

func (r *NetReader) ReadBytes(n int) ([]byte, error) {
	if n > len(r.buf) {
		return nil, ErrOutOfRange
	} else {
		out := r.buf[0:n]
		r.buf = r.buf[n:]
		r.offset += n
		return out, nil
	}
}

func (r *NetReader) Flush() []byte {
	out := r.buf
	r.buf = r.buf[:0]
	return out
}

type NetWriter struct {
	netRW_Info
	buf *bytes.Buffer
}

func (w *NetWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func (w *NetWriter) Write(data []byte) (int, error) {
	n, err := w.buf.Write(data)
	if err != nil {
		return n, err
	}
	w.offset += n
	return n, nil
}

func (w *NetWriter) Close() error {
	return nil
}

func (w *NetWriter) WriteString(s string) error {
	if err := GenericNetWriter[uint32](w, uint32(len(s))); err != nil {
		return err
	}
	_, err := w.Write([]byte(s))
	return err
}

func GenericNetReader[T NetRW_Typ](r *NetReader) (T, error) {
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
	r.offset += int(n)
	return out, nil
}

func GenericNetWriter[T NetRW_Typ](w *NetWriter, t T) (err error) {
	if err = binary.Write(w.buf, binary.BigEndian, t); err != nil {
		return
	}
	n := reflect.TypeOf(t).Size()
	w.offset += int(n)
	return nil
}
