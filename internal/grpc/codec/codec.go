package codec

import (
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
	"reflect"
)

var (
	MarshalOptions   = proto.MarshalOptions{}
	UnmarshalOptions = proto.UnmarshalOptions{}
)

type codec struct {
	name string
}

func (c *codec) Marshal(v interface{}) ([]byte, error) {
	var (
		ok bool
		df z.DataFrameInterface
	)
	if df, ok = v.(z.DataFrameInterface); !ok {
		return nil, nil
	}
	return df.Compose()
}

func (c *codec) Unmarshal(wire []byte, v interface{}) (err error) {
	var (
		df z.DataFrameInterface
	)
	if df, err = NewDataFrame(v); err != nil {
		return
	}
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(df).Elem())
	return df.Parse(wire)
}

func (c *codec) Name() string {
	return c.name
}
