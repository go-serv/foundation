package net

import (
	"context"
	meta_net "github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"
)

func refHelper(payload interface{}) (mRef z.MethodReflectionInterface, msgRef z.MessageReflectionInterface, err error) {
	msg, ok := payload.(proto.Message)
	if !ok {
		return nil, nil, nil
	}
	reflect := runtime.Runtime().Reflection()
	mRef, err = reflect.MethodReflectionFromMessage(msg)
	if err != nil {
		return
	}
	msgRef = mRef.FromMessage(msg)
	return
}

func NewClientRequest(payload interface{}, m z.MetaInterface, info *clientInfo) (r *clientRequest, err error) {
	r = new(clientRequest)
	r.payload = payload
	r.meta = m
	r.clientInfo = info
	r.methodReflect, r.msgReflect, err = refHelper(payload)
	if err != nil {
		return
	}
	return
}

func NewRequest(payload interface{}, md metadata.MD, info *clientInfo) (r *serverRequest, err error) {
	var (
		ok bool
	)
	r = new(serverRequest)
	r.payload = payload
	r.clientInfo = info
	msg, ok := payload.(proto.Message)
	if !ok {
		return nil, nil
	}
	// Meta
	r.meta = meta_net.NewMeta(md)
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
