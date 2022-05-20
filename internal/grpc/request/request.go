package request

import (
	"github.com/go-serv/service/internal/ancillary"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type request struct {
	context.Context
	meta metadata.MD
	data interface{}
}

type requestData struct {
	proto.Message
	data interface{}
}

type unaryRequest struct {
	request
}

type unaryNetRequest struct {
	unaryRequest
}

func (r *request) Data() interface{} {
	return r.data
}

func (r *request) MethodName() string {
	name, _ := grpc.Method(r.Context)
	return ancillary.GrpcDotNotation(name).MethodName()
}

func (r *request) WithData(data interface{}) {
	r.data = data
}
