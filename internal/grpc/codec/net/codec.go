package codec

import (
	"github.com/go-serv/service/internal/service"
	"google.golang.org/protobuf/proto"
)

const Name = "net-service"

type codec struct {
	svc service.NetworkServiceInterface
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return nil
}

func (codec) Name() string {
	return Name
}
