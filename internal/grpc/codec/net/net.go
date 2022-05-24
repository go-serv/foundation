package codec

import (
	cc "github.com/go-serv/service/internal/grpc/codec"
)

// Payload flags
type PacketFlags uint32

const (
	Encryption PacketFlags = 1 << iota
)

const Name = "net-service"

type codec struct {
}

type marshaler struct {
	cc.MarshalerInterface
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(cc.MarshalerInterface)
	if !ok {
		return nil, nil
	} else {
		return m.Run()
	}
}

func (c codec) Unmarshal(data []byte, v interface{}) error {
	u, err := NewUnmarshaler(data, v)
	if err != nil {
		return err
	} else {
		return u.Run()
	}
}

func (codec) Name() string {
	return Name
}
