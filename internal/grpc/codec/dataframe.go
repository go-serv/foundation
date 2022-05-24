package codec

import (
	"bytes"
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const magicWordLen = 8

type HeaderFlags32Type uint32

func (f HeaderFlags32Type) Has(chkFlag HeaderFlags32Type) bool {
	return f&chkFlag != 0
}

var (
	// head -c 8 /dev/urandom | hexdump -C
	dfMagicWord             = [magicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	errorHeaderParserFailed = status.Error(codes.Internal, "failed to parse data frame header")
)

type DataFrame struct {
	hdrFlags HeaderFlags32Type
	payload  []byte
	msg      proto.Message
	netw     *ancillary.NetWriter
	netr     *ancillary.NetReader
}

func DataFrameBuilder(b []byte) (*DataFrame, error) {
	df := new(DataFrame)
	if err := df.ParseHeader(b); err != nil {
		return nil, err
	} else {
		return df, nil
	}
}

func (df *DataFrame) ParseHeader(b []byte) error {
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
	flags, err := ancillary.GenericNetReader[uint32](df.netr)
	if err != nil {
		return errorHeaderParserFailed
	}
	df.hdrFlags = HeaderFlags32Type(flags)
	return nil
}

func (df *DataFrame) HeaderFlags() HeaderFlags32Type {
	return df.hdrFlags
}

func (df *DataFrame) WithHeaderFlag(f HeaderFlags32Type) {
	df.hdrFlags |= f
}

func (df *DataFrame) AttachData(b []byte) error {
	var err error
	if df.netw.Offset() == 0 {
		if _, err = df.netw.Write(dfMagicWord[:]); err != nil {
			return err
		}
		if err = ancillary.GenericNetWriter[uint32](df.netw, uint32(df.hdrFlags)); err != nil {
			return err
		}
	}
	if _, err := df.netw.Write(b); err != nil {
		return err
	}
	return nil
}

func (df *DataFrame) Compose() []byte {
	return df.netw.Bytes()
}

func (df *DataFrame) Payload() []byte {
	if df.netr == nil {
		panic("trying to get payload from an uninitialized data frame")
	} else {
		return df.netr.Flush()
	}
}
