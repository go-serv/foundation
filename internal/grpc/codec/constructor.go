package codec

import (
	i "github.com/go-serv/service/internal"
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
	c.clientProc = newProcessor(c)
	c.serviceProc = newProcessor(c)
	return c
}

func newProcessor(c i.CodecInterface) *msgproc {
	p := &msgproc{}
	p.codec = c
	return p
}
