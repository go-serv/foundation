package net

import "bytes"

func NewReader(data []byte) *NetReader {
	r := new(NetReader)
	r.buf = data
	return r
}

func NewWriter() *NetWriter {
	w := new(NetWriter)
	w.buf = new(bytes.Buffer)
	return w
}
