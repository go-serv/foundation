package response

import (
	"github.com/go-serv/service/internal/grpc/codec"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func NewResponse(v interface{}, md *metadata.MD) (res *response, err error) {
	res = new(response)
	if v != nil {
		res.df = codec.NewDataFrame(v.(proto.Message))
		if err = res.Populate(v.(proto.Message)); err != nil {
			return
		}
	}
	res.meta = net.NewMeta(md)
	return
}
