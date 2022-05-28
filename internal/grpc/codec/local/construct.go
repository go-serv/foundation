package net

import (
	i "github.com/go-serv/service/internal"
	cc "github.com/go-serv/service/internal/grpc/codec"
	"google.golang.org/grpc/encoding"
)

func NewOrRegistered(name string) i.CodecInterface {
	c := encoding.GetCodec(name)
	if c != nil {
		return c.(i.CodecInterface)
	} else {
		c := new(codec)
		c.CodecInterface = cc.NewCodec(name)
		encoding.RegisterCodec(c)
		return c
	}
}
