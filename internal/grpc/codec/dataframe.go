package codec

import (
	"bytes"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const magicWordLen = 8

var (
	// head -c 8 /dev/urandom | hexdump -C
	dfMagicWord             = [magicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	errorHeaderParserFailed = status.Error(codes.Internal, "failed to parse data frame header")
)

type dataFrame struct {
	hdrFlags i.HeaderFlags32Type
	payload  []byte
	netw     *ancillary.NetWriter
	netr     *ancillary.NetReader
}

func (df *dataFrame) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (df *dataFrame) ParseHook(*ancillary.NetReader) error {
	return nil
}

func (df *dataFrame) Parse(b []byte, hookFn func(netr *ancillary.NetReader) error) error {
	// Check for header magic word
	{
		df.netr = ancillary.NewNetReader(b)
		hdr, err := df.netr.ReadBytes(magicWordLen)
		if err != nil {
			return errorHeaderParserFailed
		}
		if bytes.Compare(hdr, dfMagicWord[:]) != 0 {
			return errorHeaderParserFailed
		}
	}
	// Store 32-bit header flags
	{
		flags, err := ancillary.GenericNetReader[uint32](df.netr)
		if err != nil {
			return errorHeaderParserFailed
		}
		df.hdrFlags = i.HeaderFlags32Type(flags)
	}
	// Call parser hook
	if hookFn != nil {
		if err := hookFn(df.netr); err != nil {
			return err
		}
	}
	df.payload = df.netr.Flush()
	return nil
}

func (df *dataFrame) HeaderFlags() i.HeaderFlags32Type {
	return df.hdrFlags
}

func (df *dataFrame) WithHeaderFlag(f i.HeaderFlags32Type) {
	df.hdrFlags |= f
}

// AttachData attaches data to the data frame payload.
func (df *dataFrame) AttachData(in []byte) {
	l1 := len(df.payload)
	l2 := len(in)
	buf := make([]byte, l1+l2)
	_ = copy(buf, df.payload)
	_ = copy(buf[l1:], in)
	df.payload = buf
}

func (df *dataFrame) Compose(header []byte) (out []byte, err error) {
	// Magic word
	if _, err = df.netw.Write(dfMagicWord[:]); err != nil {
		return
	}
	// Header
	if err = ancillary.GenericNetWriter[uint32](df.netw, uint32(df.hdrFlags)); err != nil {
		return
	}
	// Write header data
	df.netw.Write(header)
	// Payload
	if _, err = df.netw.Write(df.payload); err != nil {
		return
	}
	out = df.netw.Bytes()
	return
}

func (df *dataFrame) Payload() []byte {
	return df.payload
}

func (df *dataFrame) WithPayload(b []byte) {
	df.payload = b
}
