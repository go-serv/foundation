package codec

import (
	"bytes"
	"errors"
	"github.com/go-serv/service/internal/ancillary/net"
	"github.com/go-serv/service/pkg/z/ancillary/crypto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"hash/fnv"
	"math/rand"
	"reflect"
	"runtime"
	"sync/atomic"
	"unsafe"
)

const magicWordLen = 8

var (
	// head -c 8 /dev/urandom | hexdump -C
	dfMagicWord             = [magicWordLen]byte{0xd6, 0x2f, 0x7b, 0x92, 0x24, 0xfb, 0x37, 0x9c}
	errorHeaderParserFailed = status.Error(codes.Internal, "failed to parse data frame header")
	dataFramePtrPool        dataFramePtrPoolTyp
)

const ptrPoolSize = 1000

type (
	dataFramePtrPoolTyp  []dataFramePtrPoolItem
	dataFramePtrPoolItem struct {
		df    *dataFrame
		inuse uint32
	}
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

func LoadDataFrameFromPtrPool(msg proto.Message) (df *dataFrame, err error) {
	ptr1 := (*reflect.SliceHeader)(unsafe.Pointer(reflect.ValueOf(msg).Elem().FieldByName("unknownFields").Addr().Pointer()))
	idx := uint(ptr1.Data)
	if df = dataFramePtrPool[idx].df; df == nil {
		return nil, errors.New("")
	}
	return dataFramePtrPool[idx].df, nil
}

func (df dataFrame) ptrPoolIndex(msg proto.Message) (idx uint, err error) {
	protoPtr := uint64(reflect.ValueOf(msg).Elem().Field(0).Addr().Pointer())
	nw := net.NewWriter()
	if err = net.GenericNetWriter[uint64](nw, protoPtr); err != nil {
		return
	}
	fnv := fnv.New64a()
	if _, err = fnv.Write(nw.Bytes()); err != nil {
		return
	}
	hash := fnv.Sum64()
	return uint(hash % ptrPoolSize), nil
}

func (df *dataFrame) addToPtrPool() (err error) {
	//var idx uint
	//if idx, err = df.ptrPoolIndex(df.ProtoMessage()); err != nil {
	//	return
	//}
	randIdx := rand.Intn(ptrPoolSize)
	ptr1 := (*reflect.SliceHeader)(unsafe.Pointer(reflect.ValueOf(df.ProtoMessage()).Elem().FieldByName("unknownFields").Addr().Pointer()))
	ptr1.Data = uintptr(randIdx)
	for {
		if atomic.CompareAndSwapUint32(&dataFramePtrPool[randIdx].inuse, 0, 1) {
			dataFramePtrPool[randIdx].df = df
			return
		} else {
			runtime.Gosched()
		}
	}
}

func (df *dataFrame) RemoveFromPtrPool() (err error) {
	var idx uint
	if idx, err = df.ptrPoolIndex(df.ProtoMessage()); err != nil {
		return
	}
	atomic.StoreUint32(&dataFramePtrPool[idx].inuse, 0)
	return
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

func (df *dataFrame) ProtoMessage() proto.Message {
	return df.Message
}

func (df *dataFrame) WithProtoMessage(msg proto.Message) {
	df.Message = msg
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
