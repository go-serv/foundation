package net

import (
	cc "github.com/go-serv/foundation/internal/grpc/codec"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/encoding"
)

func NewOrRegistered(name string) z.CodecInterface {
	c := encoding.GetCodec(name)
	if c != nil {
		return c.(z.CodecInterface)
	} else {
		c := new(codec)
		c.CodecInterface = cc.NewCodec(name)
		encoding.RegisterCodec(c)
		return c
	}
}
