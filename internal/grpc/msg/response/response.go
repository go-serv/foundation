package response

import (
	"github.com/go-serv/foundation/internal/grpc/msg"
	"github.com/go-serv/foundation/pkg/z"
)

type response struct {
	msg.Reflection
	data any
	meta z.MetaInterface
}

func (r *response) Data() any {
	return r.data
}

func (r *response) WithData(data any) {
	r.data = data
}

func (r *response) Meta() z.MetaInterface {
	return r.meta
}
