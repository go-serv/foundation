package service

import (
	"github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/internal/reflect"
	"github.com/mesh-master/foundation/pkg/z"
)

var ref z.ReflectInterface

func init() {
	ref = reflect.NewReflection()
	ref.AddProtoExtension(foundation.E_AuthType)
	ref.AddProtoExtension(foundation.E_EnforceSecureChannel)
	ref.AddProtoExtension(foundation.E_RequiresRole)

	ref.AddProtoExtension(foundation.E_NewSession)
	ref.AddProtoExtension(foundation.E_RequireSession)
	ref.AddProtoExtension(foundation.E_CloseSession)
	ref.AddProtoExtension(foundation.E_ClientCopyMetaOff)
}
