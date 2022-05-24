package codec

import (
	cc "github.com/go-serv/service/internal/grpc/codec"
)

func NewMarshaler(msg interface{}) *marshaler {
	m := new(marshaler)
	m.MarshalerInterface = cc.NewMarshaler(msg)
	return m
}

func NewUnmarshaler(data []byte, v interface{}) (cc.UnmarshalerInterface, error) {
	return cc.NewUnmarshaler(data, v)
}
