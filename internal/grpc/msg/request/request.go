package request

import (
	"github.com/go-serv/service/pkg/z"
	"net"
)

type clientInfo struct {
	addr net.Addr
}

type request struct {
	df            z.DataFrameInterface
	meta          z.MetaInterface
	methodReflect z.MethodReflectionInterface
	msgReflect    z.MessageReflectionInterface
	clientInfo    *clientInfo
}

type serverRequest struct {
	request
}

type clientRequest struct {
	request
}

func (r *request) Meta() z.MetaInterface {
	return r.meta
}

func (r *request) MethodReflection() z.MethodReflectionInterface {
	return r.methodReflect
}

func (r *request) MessageReflection() z.MessageReflectionInterface {
	return r.msgReflect
}

func (r *request) DataFrame() z.DataFrameInterface {
	return r.df
}
