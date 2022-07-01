package response

import (
	"github.com/go-serv/service/internal/grpc/msg"
	"github.com/go-serv/service/pkg/z"
)

type response struct {
	msg.Reflection
	df   z.DataFrameInterface
	meta z.MetaInterface
}

func (r *response) DataFrame() z.DataFrameInterface {
	return r.df
}

func (r *response) WithDataFrame(df z.DataFrameInterface) {
	r.df = df
}

func (r *response) Meta() z.MetaInterface {
	return r.meta
}
