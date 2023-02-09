package codec

import (
	"github.com/mesh-master/foundation/addon/sec_chan/x"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

var (
	errMustImplementInterface = status.Errorf(codes.Internal, "codec: message must implement '%s'",
		reflect.TypeOf((*x.DataFrameInterface)(nil)).Elem().Name())
)

// Name codec name. It will be passed in the :content-type request header as content subtype.
const Name = "gs-proto-enc"

type codec struct{}

func (c codec) Marshal(v interface{}) ([]byte, error) {
	var (
		ok bool
		df x.DataFrameInterface
	)
	if df, ok = v.(x.DataFrameInterface); !ok {
		return nil, errMustImplementInterface
	}
	return df.Compose()
}

func (c codec) Unmarshal(wire []byte, v interface{}) (err error) {
	var (
		ok bool
		df x.DataFrameInterface
	)
	if df, ok = v.(x.DataFrameInterface); !ok {
		return errMustImplementInterface
	}
	if err = df.Parse(wire); err != nil {
		return
	}
	return
}

func (c codec) Name() string {
	return Name
}
