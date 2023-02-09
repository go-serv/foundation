package y

import (
	"github.com/mesh-master/foundation/addon/sec_chan/internal/codec"
	"google.golang.org/grpc/encoding"
)

func NewCodec() encoding.Codec {
	return codec.NewCodec()
}
