package request

import (
	i "github.com/go-serv/service/internal"
	"golang.org/x/net/context"
)

type Request struct {
	context.Context
	Method string
	Meta   i.MetaInterface
	Idata  interface{}
}

func (r *Request) MethodName() string {
	return r.Method
}

func (r *Request) Data() interface{} {
	return r.Idata
}

func (r *Request) WithData(data interface{}) {
	//r.data = data
}
