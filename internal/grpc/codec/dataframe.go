package codec

import (
	"bytes"
	"github.com/go-serv/foundation/internal/ancillary/net"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

const magicWordLen = 8

var (
	// head -c 8 /dev/urandom | hexdump -C
	dfMagicWord             = [magicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	errorHeaderParserFailed = status.Error(codes.Internal, "failed to parse data frame header")
)

type (
	protoMsgType uint8
)

const (
	EncryptedMessage protoMsgType = iota + 1
)

type dataFrame struct {
	proto.Message
	cipher          crypto.AEAD_CipherInterface
	hdrMsgType      protoMsgType
	hdrFlags        uint8
	hdrReserved8_A  uint8
	hdrReserved16_B uint16
	hdrReserved24_C uint32
	payload         []byte
	netw            *net.NetWriter
	netr            *net.NetReader
}

func MessageWrapperHandler() encoding.MessageWrapperHandler {
	return func(v interface{}) encoding.MessageWrapper {
		df := NewDataFrame(v.(proto.Message))
		return df
	}
}

func (df *dataFrame) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (df *dataFrame) Unmarshal(data []byte) error {
	return UnmarshalOptions.Unmarshal(data, df.Message)
}

func (df *dataFrame) Marshal() ([]byte, error) {
	return MarshalOptions.Marshal(df.Message)
}

func (df *dataFrame) Parse(wire []byte) (err error) {
	var (
		mw     []byte
		header uint64
	)
	df.netr = net.NewReader(wire)
	// Check for the data frame magic word. If there is no such, then we have an ordinary proto message.
	if mw, err = df.netr.ReadBytes(magicWordLen); err != nil {
		return df.Unmarshal(wire)
	}
	if bytes.Compare(mw, dfMagicWord[:]) != 0 {
		return df.Unmarshal(wire)
	}
	//
	if header, err = net.GenericNetReader[uint64](df.netr); err != nil {
		return err
	}
	df.hdrMsgType = protoMsgType(header & 0xff)
	df.hdrFlags = uint8(header >> 8)
	df.payload = df.netr.Flush()
	return
}

func (df *dataFrame) packHeader() uint64 {
	header := uint64(df.hdrReserved24_C)<<40 |
		uint64(df.hdrReserved16_B)<<24 |
		uint64(df.hdrReserved8_A)<<16 |
		uint64(df.hdrFlags)<<8 |
		uint64(df.hdrMsgType)
	return header
}

func (df *dataFrame) unpackHeader(hdr uint64) {
	df.hdrMsgType = protoMsgType(hdr)
	df.hdrFlags = uint8(hdr >> 8)
	df.hdrReserved8_A = uint8(hdr >> 16)
	df.hdrReserved16_B = uint16(hdr >> 24)
	df.hdrReserved24_C = uint32(hdr >> 40)
}

func (df *dataFrame) headerBytes(hdr uint64) (b []byte, err error) {
	w := net.NewWriter()
	if err = net.GenericNetWriter[uint64](w, hdr); err != nil {
		return
	}
	return w.Bytes(), nil
}

func (df *dataFrame) HeaderFlags() uint8 {
	return df.hdrFlags
}

func (df *dataFrame) WithHeaderFlag(f uint8) {
	df.hdrFlags |= f
}

func (df *dataFrame) Compose() (out []byte, err error) {
	if df.hdrMsgType == 0 {
		return df.Marshal()
	}
	// Write the data frame magic word.
	if _, err = df.netw.Write(dfMagicWord[:]); err != nil {
		return
	}
	// Write data frame header.
	header := df.packHeader()
	if err = net.GenericNetWriter[uint64](df.netw, header); err != nil {
		return
	}
	if out, err = MarshalOptions.Marshal(df.Message); err != nil {
		return
	}
	// Encrypt payload.
	if df.cipher != nil {
		var hdrBytes []byte
		if hdrBytes, err = df.headerBytes(header); err != nil {
			return
		}
		out = df.cipher.Encrypt(out, hdrBytes)
	}
	df.netw.Write(out)
	out = df.netw.Bytes()
	return
}

func (df *dataFrame) Interface() interface{} {
	return df.Message
}

func (df *dataFrame) WithBlockCipher(cipher crypto.AEAD_CipherInterface) {
	df.cipher = cipher
	df.hdrMsgType = EncryptedMessage
}

func (df *dataFrame) Decrypt() (err error) {
	var out, hdrBytes []byte
	if hdrBytes, err = df.headerBytes(df.packHeader()); err != nil {
		return
	}
	if out, err = df.cipher.Decrypt(df.payload, hdrBytes); err != nil {
		return
	}
	err = df.Unmarshal(out)
	return
}

func (df *dataFrame) Payload() []byte {
	return df.payload
}
