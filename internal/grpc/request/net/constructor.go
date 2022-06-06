package net

import (
	"context"
	meta_net "github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/runtime"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

func NewRequest(payload interface{}, md metadata.MD, info *clientInfo) (r *request, err error) {
	var (
		ok bool
	)
	r = new(request)
	r.payload = payload
	r.clientInfo = info
	msg, ok := payload.(proto.Message)
	if !ok {
		return nil, nil
	}
	// Meta
	r.meta = meta_net.NewServerMeta(md)
	err = r.meta.Hydrate()
	if err != nil {
		return
	}
	// Reflection
	reflect := runtime.Runtime().Reflection()
	r.methodReflect, err = reflect.MethodReflectionFromMessage(msg)
	if err != nil {
		return
	}
	r.msgReflect = r.methodReflect.FromMessage(msg)
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
