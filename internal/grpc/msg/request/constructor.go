package request

import (
	"context"
	"github.com/go-serv/foundation/internal/grpc"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

func NewClientRequest(in interface{}, m z.MetaInterface, info *clientInfo) (r *clientRequest, err error) {
	var (
		ok  bool
		msg proto.Message
	)
	r = new(clientRequest)
	r.data = in
	r.meta = m
	r.clientInfo = info
	if msg, ok = in.(proto.Message); !ok {
		return nil, grpc.ErrInvalidProtoMessage
	}
	err = r.Populate(msg)
	return
}

func NewServerRequest(data interface{}, md *metadata.MD, info *clientInfo) (r *serverRequest, err error) {
	var (
		ok  bool
		msg proto.Message
	)
	r = new(serverRequest)
	r.data = data
	r.clientInfo = info
	r.meta = meta_net.NewRequestMeta(md)
	if msg, ok = data.(proto.Message); !ok {
		return nil, grpc.ErrInvalidProtoMessage
	}
	if err = r.Populate(msg); err != nil {
		return
	}
	if err = r.meta.Hydrate(); err != nil {
		return
	}
	return
}

func NewClientInfo(ctx context.Context) *clientInfo {
	info := new(clientInfo)
	client, ok := peer.FromContext(ctx)
	if ok {
		info.addr = client.Addr
	}
	return info
}
