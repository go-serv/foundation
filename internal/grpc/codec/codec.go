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
		df *dataFrame
	)
	if df, err = NewDataFrame(v); err != nil {
		return
	}
	//v.(proto.Message).ProtoReflect().SetUnknown(protoreflect.RawFields{})
	// Map the incoming proto message to its data frame.
	if err = df.Parse(wire); err != nil {
		return
	}
	if err = df.addToPtrPool(); err != nil {
		return
	}
	//ptr1 := (*reflect.SliceHeader)(unsafe.Pointer(reflect.ValueOf(v).Elem().FieldByName("unknownFields").Addr().Pointer()))
	//ptr1.Data = uintptr(reflect.ValueOf(df).Elem().UnsafePointer())
	return
}

func (c *codec) Name() string {
	return c.name
}
