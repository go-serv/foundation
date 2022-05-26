package codec

import (
	cc "github.com/go-serv/service/internal/grpc/codec"
)

func newCodec() cc.CodecInterface {
	c := new(codec)
	c.CodecInterface = cc.NewCodec(Name)
	return c
}
