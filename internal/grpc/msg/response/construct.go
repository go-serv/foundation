package response

import (
	"github.com/go-serv/foundation/internal/grpc"
	"github.com/go-serv/foundation/internal/grpc/meta/net"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func NewResponse(data interface{}, md *metadata.MD) (res *response, err error) {
	var (
		ok  bool
		msg proto.Message
	)
	res = new(response)
	if data != nil {
		if msg, ok = data.(proto.Message); !ok {
			return nil, grpc.ErrInvalidProtoMessage
		}
		if err = res.Populate(msg); err != nil {
			return
		}
		res.data = data
	}
	res.meta = net.NewMeta(md)
	return
}
