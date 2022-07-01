package request

import (
	"context"
	meta_net "github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

func NewClientRequest(df z.DataFrameInterface, m z.MetaInterface, info *clientInfo) (r *clientRequest, err error) {
	r = new(clientRequest)
	r.df = df
	r.meta = m
	r.clientInfo = info
	reflect := runtime.Runtime().Reflection()
	if r.methodReflect, err = reflect.MethodReflectionFromMessage(df.Interface().(proto.Message)); err != nil {
		return
	}
	r.msgReflect = r.methodReflect.FromMessage(df.Interface().(proto.Message))
	return
}

func NewRequest(df z.DataFrameInterface, md *metadata.MD, info *clientInfo) (r *serverRequest, err error) {
	r = new(serverRequest)
	r.df = df
	r.clientInfo = info
	// Meta
	r.meta = meta_net.NewMeta(md)
	if err = r.meta.Hydrate(); err != nil {
		return
	}
	// Reflection
	reflect := runtime.Runtime().Reflection()
	if r.methodReflect, err = reflect.MethodReflectionFromMessage(df.Interface().(proto.Message)); err != nil {
		return
	}
	r.msgReflect = r.methodReflect.FromMessage(df.Interface().(proto.Message))
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
