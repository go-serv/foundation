package codec

import (
	"github.com/go-serv/service/internal/ancillary"
	"google.golang.org/protobuf/proto"
)

func NewDataFrame() *DataFrame {
	df := new(DataFrame)
	return df
}

func NewCodec() *codec {
	c := &codec{}
	c.interceptorsChain = make([]CodecInterceptorHandler, 0)
	c.df = NewDataFrame()
	c.df.netw = ancillary.NewNetWriter()
	c.msgUnmarshaler = make(msgUnmarshalerMap, 10000)
	return c
}

func NewMarshaler(msg interface{}) *marshaler {
	m := new(marshaler)
	m.Message = msg.(proto.Message)
	m.codec = NewCodec()
	m.codec.ChainInterceptorHandler(m.marshal)
	return m
}

func NewUnmarshaler(b []byte, v interface{}) (*unmarshaler, error) {
	var err error
	un := new(unmarshaler)
	un.v = v
	un.codec = NewCodec()
	un.codec.df, err = DataFrameBuilder(b)
	if err != nil {
		return un, err
	}
	return un, nil
}
