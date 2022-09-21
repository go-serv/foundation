package codec

import (
	net_io "github.com/go-serv/foundation/pkg/ancillary/net/io"
	"google.golang.org/protobuf/proto"
)

func NewDataFrame(msg proto.Message) *dataFrame {
	df := new(dataFrame)
	df.Message = msg
	df.netw = net_io.NewWriter()
	return df
}

func NewCodec() codec {
	c := codec{}
	return c
}
