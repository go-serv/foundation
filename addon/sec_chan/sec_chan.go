package sec_chan

import (
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"github.com/go-serv/foundation/pkg/z"
)

func NewCodec() z.CodecInterface {
	return codec.NewCodec()
}
