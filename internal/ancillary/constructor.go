package ancillary

import "bytes"

func NewNetReader(data []byte) *NetReader {
	r := new(NetReader)
	r.buf = data
	return r
}

func NewNetWriter() *NetWriter {
	w := new(NetWriter)
	w.buf = new(bytes.Buffer)
	return w
}
