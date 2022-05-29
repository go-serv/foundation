package local

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	local_cc "github.com/go-serv/service/internal/grpc/codec/local"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewService(name protoreflect.FullName) *localService {
	s := &localService{service.NewBaseService(name)}
	s.ServiceInterface = service.NewBaseService(name)
	cc := local_cc.NewOrRegistered(string(name))
	s.ServiceInterface.WithCodec(cc)
	//
	sd := s.ServiceInterface.Descriptor()
	sd.AddMessageExtension(go_serv.E_LocalShmIpc)
	sd.Populate()
	//
	runtime.Runtime().RegisterLocalService(s)
	return s
}
