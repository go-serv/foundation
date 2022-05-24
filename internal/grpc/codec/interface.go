package codec

type DataFrameInterface interface {
	HeaderFlags() HeaderFlags32Type
	WithHeaderFlag(HeaderFlags32Type)
	Compose() []byte
	Payload() []byte
	AttachData(b []byte) error
}

type CodecInterface interface {
	ChainInterceptorHandler(CodecInterceptorHandler)
}

type MarshalerInterface interface {
	DataFrame() DataFrameInterface
	CodecInterface
	Run() ([]byte, error)
}
type DoNotImplement interface{ ProtoInternal(DoNotImplement) }

type UnmarshalerInterface interface {
	DataFrame() DataFrameInterface
	CodecInterface
	Run() error
	DoNotImplement
}
