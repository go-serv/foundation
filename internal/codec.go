package internal

import (
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

type HeaderFlags32Type uint32

func (f HeaderFlags32Type) Has(chkFlag HeaderFlags32Type) bool {
	return f&chkFlag != 0
}

type DataFrameInterface interface {
	Parse([]byte, func(netr *ancillary.NetReader) error) error
	ParseHook(*ancillary.NetReader) error
	HeaderFlags() HeaderFlags32Type
	WithHeaderFlag(HeaderFlags32Type)
	Compose(header []byte) ([]byte, error)
	Payload() []byte
	WithPayload([]byte)
	AttachData(b []byte)
}

type LocalDataFrameInterface interface {
	DataFrameInterface
	SharedMemObjectName() string
	WithSharedMemObjectName(string)
	SharedMemBlockSize() uint32
	WithSharedMemBlockSize(uint32)
	SharedMemDataSize() uint32
	WithSharedMemDataSize(uint32)
}

type UnmarshalMwTaskHandler func(in []byte, mf MethodReflectInterface, msgRef MessageReflectInterface, df DataFrameInterface) ([]byte, error)
type MarshalMwTaskHandler func(in []byte, mf MethodReflectInterface, msgRef MessageReflectInterface, df DataFrameInterface) ([]byte, error)

type CodecMwTaskInterface interface {
	Execute() ([]byte, error)
}

type CodecMiddlewareGroupInterface interface {
	AddHandlers(UnmarshalMwTaskHandler, MarshalMwTaskHandler)
	NewUnmarshalTask(wire []byte, msg proto.Message) (CodecMwTaskInterface, error)
	NewMarshalTask(wire []byte, msg proto.Message) (CodecMwTaskInterface, error)
}

type CodecInterface interface {
	encoding.Codec
	PureUnmarshal(wire []byte, m proto.Message) error
	PureMarshal(m proto.Message) ([]byte, error)
	NewDataFrame() DataFrameInterface
}

type CodecAwareInterface interface {
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
	CodecMiddlewareGroup() CodecMiddlewareGroupInterface
}
