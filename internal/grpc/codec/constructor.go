package codec

import (
	"github.com/go-serv/service/internal/ancillary/net"
)

func NewDataFrame() *dataFrame {
	df := new(dataFrame)
	df.netw = net.NewWriter()
	return df
}

func NewCodec(name string) *codec {
	c := &codec{}
	c.name = name
	return c
}
