package ancillary

import "bytes"

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
