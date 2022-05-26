package codec

import (
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

type DataFrameInterface interface {
	Parse([]byte) error
	ParseHook() error
	HeaderFlags() HeaderFlags32Type
	WithHeaderFlag(HeaderFlags32Type)
	Compose() ([]byte, error)
	Payload() []byte
	AttachData(b []byte)
}

type CodecInterface interface {
	encoding.Codec
	Processor() MessageProcessorInterface
	NewDataFrame() DataFrameInterface
}

type MessageProcessorInterface interface {
	AddHandlers(pre TaskHandler, post TaskHandler)
	NewPreTask(wire []byte, msg proto.Message) (*msgprocPreTask, error)
	NewPostTask(wire []byte, msg proto.Message) (*msgprocPostTask, error)
}
