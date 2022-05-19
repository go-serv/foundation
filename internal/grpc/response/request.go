package response

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type response struct {
	context.Context
	meta metadata.MD
	data interface{}
}

type localResponse struct {
	response
}

type netResponse struct {
	response
}
