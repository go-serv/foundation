package codec

import (
	"github.com/go-serv/service/internal/ancillary"
)

func NewDataFrame() *dataFrame {
	df := new(dataFrame)
	df.netw = ancillary.NewNetWriter()
	return df
}

func NewCodec(name string) *codec {
	c := &codec{}
	c.name = name
	return c
}
