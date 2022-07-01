package codec

import (
	"errors"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
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
		return nil, errors.New("not a dataframe")
	}
	return df.Compose()
}

func (c *codec) Unmarshal(wire []byte, v interface{}) (err error) {
	var (
		ok bool
		df z.DataFrameInterface
	)
	if df, ok = v.(z.DataFrameInterface); !ok {
		return errors.New("expected interface z.Dataframe")
	}
	if err = df.Parse(wire); err != nil {
		return
	}
	return
}

func (c *codec) Name() string {
	return c.name
}
