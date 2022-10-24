package y

import (
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"google.golang.org/grpc/encoding"
)

func NewCodec() encoding.Codec {
	return codec.NewCodec()
}
