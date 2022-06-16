package runtime

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/platform"
	"github.com/go-serv/service/internal/reflect"
)

var rt *runtime

func init() {
	rt = new(runtime)
	rt.platform = platform.NewPlatform()
	// Add the protobuf extensions
	rt.ref = reflect.NewReflection()
	rt.ref.AddProtoExtension(go_serv.E_NewInsecureSession)
	rt.ref.AddProtoExtension(go_serv.E_RequireSession)
	rt.ref.AddProtoExtension(go_serv.E_OptionalSession)
	rt.ref.AddProtoExtension(go_serv.E_CloseSession)
	rt.ref.AddProtoExtension(go_serv.E_CopyMetaOff)
	//
	rt.resolvers = make(resolversMapTyp, 0)
	rt.eventsMap = make(eventsMapTyp, 0)
	//
	rt.netServices = make(registry, 0)
	rt.netClients = make(registry, 0)
	rt.localService = make(registry, 0)
	rt.localClients = make(registry, 0)
}
