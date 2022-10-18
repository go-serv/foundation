package sec_chan

import (
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterMessageWrapper(codec.Name, codec.MessageWrapperHandler())
	encoding.RegisterCodec(codec.NewCodec())
}
