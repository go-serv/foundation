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
	df         z.DataFrameInterface
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

func (r *request) DataFrame() z.DataFrameInterface {
	return r.df
}
