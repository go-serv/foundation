package net

import (
	"github.com/go-serv/service/internal/grpc/meta/net"
)

func NewResponse() *response {
	r := new(response)
	r.meta = net.NewClientMeta()
	return r
}
