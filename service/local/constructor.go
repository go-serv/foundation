package local

import (
	"github.com/go-serv/foundation/service"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewService(name protoreflect.FullName) *localService {
	svc := &localService{service.NewBaseService(name)}
	//svc.ServiceInterface = service.NewBaseService(name)
	////cc := local_cc.NewOrRegistered(string(name))
	////svc.WithCodec(cc)
	//mw_shmem.ServiceInit(svc)
	////
	//rt := runtime.Runtime()
	//rt.Reflection().AddService(name)
	//rt.Reflection().Populate()
	////rt.RegisterLocalService(svc)
	return svc
}
