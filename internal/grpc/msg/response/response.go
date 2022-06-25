package response

import (
	"github.com/go-serv/service/pkg/z"
)

type response struct {
	df            z.DataFrameInterface
	meta          z.MetaInterface
	methodReflect z.MethodReflectionInterface
	msgReflect    z.MessageReflectionInterface
}

func (r *response) DataFrame() z.DataFrameInterface {
	return r.df
}

func (r *response) Meta() z.MetaInterface {
	return r.meta
}

func (r *response) MethodReflection() z.MethodReflectionInterface {
	return r.methodReflect
}

func (r *response) MessageReflection() z.MessageReflectionInterface {
	return r.msgReflect
}

func (r *response) ToGrpcResponse() interface{} {
	return r.methodReflect
}
