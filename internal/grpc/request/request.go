package request

import (
	"github.com/go-serv/service/internal/ancillary"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type request struct {
	context.Context
	meta metadata.MD
	data interface{}
}

type unaryRequest struct {
	request
}

type unaryNetRequest struct {
	unaryRequest
}

func (r *request) MethodName() string {
	name, _ := grpc.Method(r.Context)
	return ancillary.GrpcDotNotation(name).MethodName()
}
