package request

import (
	"context"
	meta_net "github.com/go-serv/foundation/internal/grpc/meta/net"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

func NewClientRequest(df z.DataFrameInterface, m z.MetaInterface, info *clientInfo) (r *clientRequest, err error) {
	r = new(clientRequest)
	r.df = df
	r.meta = m
	r.clientInfo = info
	r.Populate(df.Interface().(proto.Message))
	return
}

func NewServerRequest(df z.DataFrameInterface, md *metadata.MD, info *clientInfo) (r *serverRequest, err error) {
	r = new(serverRequest)
	r.df = df
	r.clientInfo = info
	r.meta = meta_net.NewMeta(md)
	if err = r.Populate(df.Interface().(proto.Message)); err != nil {
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
