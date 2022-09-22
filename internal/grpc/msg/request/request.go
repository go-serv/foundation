package request

import (
	"github.com/go-serv/foundation/internal/grpc/msg"
	"github.com/go-serv/foundation/pkg/z"
	"net"
)

type clientInfo struct {
	addr net.Addr
}

type request struct {
	msg.Reflection
	data       any
	meta       z.MetaInterface
	clientInfo *clientInfo
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

func (r *request) Data() any {
	return r.data
}
