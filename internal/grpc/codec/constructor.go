package codec

import (
	"errors"
	"github.com/go-serv/service/internal/ancillary/net"
	"google.golang.org/protobuf/proto"
)

func NewDataFrame(v interface{}) (df *dataFrame, err error) {
	df = new(dataFrame)
	if v != nil {
		var (
			msg proto.Message
			ok  bool
		)
		if msg, ok = v.(proto.Message); !ok {
			return nil, errors.New("codec: message must implement proto.Message")
		}
		df.Message = msg
	}
	df.netw = net.NewWriter()
	return
}

func NewCodec(name string) *codec {
	c := &codec{}
	c.name = name
	return c
}
