package runtime

import (
	"github.com/go-serv/service/internal/reflect"
)

var rt *runtime

func init() {
	rt = new(runtime)
	rt.ref = reflect.NewReflection()
	//rt.ref.AddProtoExtension(go_serv.E_LocalShmIpc)
	//rt.ref.AddProtoExtension(go_serv.E_NetMsgEnc)
	//
	rt.eventsMap = make(eventsMapTyp, 0)
	//
	rt.netServices = make(registry, 0)
	rt.netClients = make(registry, 0)
	rt.localService = make(registry, 0)
	rt.localClients = make(registry, 0)
}
