package net

import (
	"github.com/go-serv/service/pkg/z"
	"net"
)

type clientInfo struct {
	addr net.Addr
}

type request struct {
	payload       interface{}
	meta          z.MetaInterface
	clientInfo    *clientInfo
	methodReflect z.MethodReflectionInterface
	msgReflect    z.MessageReflectionInterface
}

func (r *request) Meta() z.MetaInterface {
	return r.meta
}

func (r *request) Payload() interface{} {
	return r.payload
}

func (r *request) WithPayload(payload interface{}) {
	r.payload = payload
}

func (r *request) MethodReflection() z.MethodReflectionInterface {
	return nil
}

func (r *request) MessageReflection() z.MessageReflectionInterface {
	return nil
}
