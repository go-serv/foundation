package local

import (
	mw_shmem "github.com/go-serv/service/internal/middleware/codec/shm_ipc"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewService(name protoreflect.FullName) *localService {
	svc := &localService{service.NewBaseService(name)}
	svc.ServiceInterface = service.NewBaseService(name)
	//cc := local_cc.NewOrRegistered(string(name))
	//svc.WithCodec(cc)
	mw_shmem.ServiceInit(svc)
	//
	rt := runtime.Runtime()
	rt.Reflection().AddService(name)
	rt.Reflection().Populate()
	//rt.RegisterLocalService(svc)
	return svc
}
