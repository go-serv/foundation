package response

import (
	"github.com/go-serv/service/internal/grpc/codec"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"google.golang.org/grpc/metadata"
)

func NewResponse(v interface{}, md *metadata.MD) (res *response, err error) {
	res = new(response)
	if res.df, err = codec.NewDataFrame(v); err != nil {
		return
	}
	res.meta = net.NewMeta(md)
	return
}
