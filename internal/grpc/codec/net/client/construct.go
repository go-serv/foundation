package codec

import (
	i "github.com/go-serv/service/internal"
	cc "github.com/go-serv/service/internal/grpc/codec"
)

func newCodec() i.CodecInterface {
	c := new(codec)
	c.CodecInterface = cc.NewCodec(Name)
	return c
}
