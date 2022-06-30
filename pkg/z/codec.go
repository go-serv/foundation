package z

import (
	"github.com/go-serv/service/pkg/z/ancillary/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

type HeaderFlagsType uint64
type HeaderType uint64

func (f HeaderFlagsType) Has(chkFlag HeaderFlagsType) bool {
	return f&chkFlag != 0
}

type DataFrameInterface interface {
	Parse(wire []byte) error
	//HeaderFlags() HeaderFlagsType
	//WithHeaderFlag(HeaderFlagsType)
	Compose() ([]byte, error)
	ProtoMessage() proto.Message
	WithProtoMessage(msg proto.Message)
	WithBlockCipher(cipher crypto.AEAD_CipherInterface)
	Decrypt() error
	Payload() []byte
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	RemoveFromPtrPool() error
}

//type LocalDataFrameInterface interface {
//	DataFrameInterface
//	SharedMemObjectName() string
//	WithSharedMemObjectName(string)
//	SharedMemBlockSize() uint32
//	WithSharedMemBlockSize(uint32)
//	SharedMemDataSize() uint32
//	WithSharedMemDataSize(uint32)
//}

type CodecMwTaskUnHandler func(next MwChainElement, in []byte, method MethodReflectionInterface, msg MessageReflectionInterface, df DataFrameInterface) ([]byte, error)
type CodecMwTaskMarshalHandler func(next MwChainElement, in []byte, method MethodReflectionInterface, msg MessageReflectionInterface, df DataFrameInterface) ([]byte, error)
type NetMiddlewareReqHandler func(next MwChainElement, req RequestInterface, h grpc.UnaryHandler) (ResponseInterface, error)
type NetMiddlewareResHandler func(next MwChainElement, req ResponseInterface) (proto.Message, error)
type MwChainElement func(in []byte) (MwChainElement, error)

type CodecMwMarshalTaskInterface interface {
	Execute() ([]byte, error)
}

type CodecMwUnmarshalTaskInterface interface {
	Execute() error
}

type CodecMiddlewareGroupInterface interface {
	AddHandlers(CodecMwTaskUnHandler, CodecMwTaskMarshalHandler)
	NewUnmarshalTask(wire []byte, msg proto.Message) (CodecMwUnmarshalTaskInterface, error)
	NewMarshalTask(msg proto.Message) (CodecMwMarshalTaskInterface, error)
}

type CodecInterface interface {
	encoding.Codec
	//PureUnmarshal(wire []byte, m proto.Message) error
	//PureMarshal(m proto.Message) ([]byte, error)
	//NewDataFrame() DataFrameInterface
}

type CodecAwareInterface interface {
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
}
