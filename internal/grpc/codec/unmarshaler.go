package codec

import "google.golang.org/protobuf/proto"

type unmarshaler struct {
	*codec
	v    interface{}
	opts proto.UnmarshalOptions
}

// Marshaler: 	wire data -> int1 -> int2 ->  ...
// Unmarshaler: wire data: <- int1 <- int2 <- ...
// Run invokes interceptor handlers in reverse order and unmarshales wire data
func (u *unmarshaler) Run() error {
	var data []byte = nil
	var err error
	chain := u.codec.interceptorsChain
	for i := len(chain) - 1; i >= 0; i-- {
		data, err = chain[i](data)
		if err != nil {
			return err
		}
	}
	err = u.opts.Unmarshal(u.df.Payload(), u.v.(proto.Message))
	if err != nil {
		return err
	}
	u.codec.mapProtoMessage(u.v, u)
	return nil
}

func (u *unmarshaler) DataFrame() DataFrameInterface {
	return u.df
}

func (u *unmarshaler) WithOptions(opts proto.UnmarshalOptions) {
	u.opts = opts
}

func (u *unmarshaler) ProtoInternal(DoNotImplement) {
}
