package request

import (
	"github.com/mesh-master/foundation/internal/grpc/msg"
	"github.com/mesh-master/foundation/pkg/z"
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

func (r *request) WithData(data any) {
	r.data = data
}
