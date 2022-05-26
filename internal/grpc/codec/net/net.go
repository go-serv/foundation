package codec

import (
	cc "github.com/go-serv/service/internal/grpc/codec"
)

type headerFlags cc.HeaderFlags32Type

const (
	Encryption headerFlags = 1 << iota
)

const Name = "net-service"

type codec struct {
	cc.CodecInterface
}

//func (c *codec) DataFrameBuilderHook(b []byte) (cc.DataFrameInterface, error) {
//	//return cc.DataFrameBuilder(b)
//}
