package net

import (
	z "github.com/go-serv/service/internal"
)

type response struct {
	payload interface{}
	meta    z.MetaInterface
}

func (r *response) Payload() interface{} {
	return r.payload
}

func (r *response) WithPayload(payload interface{}) {
	r.payload = payload
}

func (r *response) Meta() z.MetaInterface {
	return r.meta
}

func (r *response) MethodReflection() z.MethodReflectionInterface {
	return nil
}

func (r *response) MessageReflection() z.MessageReflectionInterface {
	return nil
}

func (r *response) ToGrpcResponse() interface{} {
	return nil
}
