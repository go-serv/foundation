package sec_chan_mw

import (
	"github.com/go-serv/foundation/addon/sec-chan-mw/internal/codec"
	mw "github.com/go-serv/foundation/addon/sec-chan-mw/internal/middleware"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/encoding"
)

func MiddlewareServerInit(m z.NetMiddlewareInterface) {
	encoding.RegisterMessageWrapper(codec.Name, codec.MessageWrapperHandler())
	mw.ServerInit(m)
}

func MiddlewareClientInit(m z.NetMiddlewareInterface) {
	encoding.RegisterMessageWrapper(codec.Name, codec.MessageWrapperHandler())
	mw.ClientInit(m)
}
