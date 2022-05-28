package local

import (
	"github.com/go-serv/service/internal/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewService(name protoreflect.FullName) *localService {
	s := &localService{service.NewBaseService(name)}
	return s
}
