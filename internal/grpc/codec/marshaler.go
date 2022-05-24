package codec

import (
	"google.golang.org/protobuf/proto"
)

type marshaler struct {
	proto.Message
	*codec
}

func (m *marshaler) marshal(in []byte) ([]byte, error) {
	return proto.Marshal(m)
}

func (m *marshaler) DataFrame() DataFrameInterface {
	return m.codec.df
}

func (m *marshaler) Run() ([]byte, error) {
	var data []byte = nil
	var err error
	chain := m.codec.interceptorsChain
	for i := 0; i < len(chain); i++ {
		data, err = chain[i](data)
		if err != nil {
			return nil, err
		}
	}
	m.df.AttachData(data)
	return m.df.Compose(), nil
}
