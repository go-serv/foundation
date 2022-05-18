package request

import (
	"golang.org/x/net/context"
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
