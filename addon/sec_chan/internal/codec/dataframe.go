package codec

import (
	"bytes"
	net_io "github.com/go-serv/foundation/pkg/ancillary/net/io"
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
	hdrFlagsType uint16
)

func (b hdrFlagsType) Set(flag hdrFlagsType) hdrFlagsType    { return b | flag }
func (b hdrFlagsType) Clear(flag hdrFlagsType) hdrFlagsType  { return b &^ flag }
func (b hdrFlagsType) Toggle(flag hdrFlagsType) hdrFlagsType { return b ^ flag }
func (b hdrFlagsType) Has(flag hdrFlagsType) bool            { return b&flag != 0 }

const (
	EncryptedFlag hdrFlagsType = 1 << iota
	ApiKeyFlag
)

type dataFrame struct {
	proto.Message
	cipher          crypto.AEAD_CipherInterface
	hdrFlags        hdrFlagsType
	hdrReserved8_A  uint8
	hdrReserved16_B uint16
	hdrReserved24_C uint32
	payload         []byte
	apiKey          []byte
	netw            *net_io.NetWriter
	netr            *net_io.NetReader
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

func (df *dataFrame) Parse(wire []byte) (err error) {
	var (
		mw     []byte
		header uint64
	)
	df.netr = net_io.NewReader(wire)
	// Check for the data frame magic word. If there is no such, then we have an ordinary proto message.
	if mw, err = df.netr.ReadBytes(magicWordLen); err != nil {
		return df.unmarshal(wire)
	}
	if bytes.Compare(mw, dfMagicWord[:]) != 0 {
		return df.unmarshal(wire)
	}
	//
	if header, err = net_io.GenericNetReader[uint64](df.netr); err != nil {
		return err
	}
	df.hdrFlags = hdrFlagsType(header & 0xffff)

	if df.hdrFlags.Has(ApiKeyFlag) {
		if df.apiKey, err = df.netr.ReadLengthPrefixed(); err != nil {
			return
		}
	}

	df.payload = df.netr.Flush()
	if !df.hdrFlags.Has(EncryptedFlag) {
		err = df.unmarshal(df.payload)
	}
	return
}

func (df *dataFrame) packHeader() uint64 {
	header := uint64(df.hdrReserved24_C)<<40 |
		uint64(df.hdrReserved16_B)<<24 |
		uint64(df.hdrReserved8_A)<<16 |
		uint64(df.hdrFlags)
	return header
}

func (df *dataFrame) unpackHeader(hdr uint64) {
	df.hdrFlags = hdrFlagsType(hdr & 0xffff)
	df.hdrReserved8_A = uint8(hdr >> 16)
	df.hdrReserved16_B = uint16(hdr >> 24)
	df.hdrReserved24_C = uint32(hdr >> 40)
}

func (df *dataFrame) headerBytes(hdr uint64) (b []byte, err error) {
	w := net_io.NewWriter()
	if err = net_io.GenericNetWriter[uint64](w, hdr); err != nil {
		return
	}
	return w.Bytes(), nil
}

func (df *dataFrame) Compose() (out []byte, err error) {
	//if df.hdrFlags == 0 {
	//	fmt.Println("marshal no enc")
	//	return df.Marshal()
	//}
	// Write the data frame magic word.
	if _, err = df.netw.Write(dfMagicWord[:]); err != nil {
		return
	}
	// Write data frame header.
	header := df.packHeader()
	if err = net_io.GenericNetWriter[uint64](df.netw, header); err != nil {
		return
	}

	if df.hdrFlags.Has(ApiKeyFlag) {
		if err = df.netw.WriteLengthPrefixed(df.apiKey); err != nil {
			return
		}
	}

	var msg []byte
	if msg, err = df.marshal(); err != nil {
		return
	}
	if _, err = df.netw.Write(msg); err != nil {
		return
	}

	kl := len(df.ApiKey())
	payload := make([]byte, kl+len(msg))
	copy(payload, df.ApiKey())
	copy(payload[kl:], msg)

	if df.cipher != nil { // Encrypt payload.
		var hdrBytes []byte
		if hdrBytes, err = df.headerBytes(header); err != nil {
			return
		}
		payload = df.cipher.Encrypt(payload, hdrBytes)
	}
	if _, err = df.netw.Write(payload); err != nil {
		return
	}
	out = df.netw.Bytes()
	return
}

func (df *dataFrame) Interface() interface{} {
	return df.Message
}

func (df *dataFrame) WithBlockCipher(cipher crypto.AEAD_CipherInterface) {
	df.cipher = cipher
	df.hdrFlags = df.hdrFlags.Set(EncryptedFlag)
}

func (df *dataFrame) Decrypt() (err error) {
	var out, hdrBytes []byte
	if hdrBytes, err = df.headerBytes(df.packHeader()); err != nil {
		return
	}
	if out, err = df.cipher.Decrypt(df.payload, hdrBytes); err != nil {
		return
	}
	err = df.unmarshal(out)
	return
}

func (df *dataFrame) Payload() []byte {
	return df.payload
}

func (df *dataFrame) ApiKey() []byte {
	return df.apiKey
}

func (df *dataFrame) WithApiKey(key []byte) {
	df.apiKey = key
	df.hdrFlags = df.hdrFlags.Set(ApiKeyFlag)
}

func (df *dataFrame) unmarshal(data []byte) error {
	return proto.Unmarshal(data, df.Message)
}

func (df *dataFrame) marshal() ([]byte, error) {
	return proto.Marshal(df.Message)
}
