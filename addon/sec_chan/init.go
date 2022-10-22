package sec_chan

import (
	"github.com/go-serv/foundation/addon/sec_chan/internal/codec"
	"github.com/go-serv/foundation/internal/autogen/net/sec_chan"
	"github.com/go-serv/foundation/internal/service"
	"google.golang.org/grpc/encoding"
)

func init() {
	encoding.RegisterMessageWrapper(codec.Name, codec.MessageWrapperHandler())
	encoding.RegisterCodec(codec.NewCodec())
	service.Reflection().AddProtoExtension(sec_chan.E_EncOff)
}
