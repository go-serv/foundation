package internal

import (
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

type HeaderFlags32Type uint32

func (f HeaderFlags32Type) Has(chkFlag HeaderFlags32Type) bool {
	return f&chkFlag != 0
}

type DataFrameInterface interface {
	Parse([]byte) error
	ParseHook() error
	HeaderFlags() HeaderFlags32Type
	WithHeaderFlag(HeaderFlags32Type)
	Compose() ([]byte, error)
	Payload() []byte
	WithPayload([]byte)
	AttachData(b []byte)
}

type UnmarshalMwTaskHandler func(in []byte, md MethodDescriptorInterface, df DataFrameInterface) ([]byte, error)
type MarshalMwTaskHandler func(in []byte, md MethodDescriptorInterface, df DataFrameInterface) ([]byte, error)

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
	NewDataFrame() DataFrameInterface
}
