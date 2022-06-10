package net

import (
	"github.com/go-serv/service/internal/grpc/meta/net"
	"google.golang.org/grpc/metadata"
)

func NewResponse(payload interface{}, md *metadata.MD) *response {
	r := new(response)
	r.payload = payload
	r.meta = net.NewMeta(md)
	return r
}
