package service

import (
	"github.com/go-serv/foundation/internal/autogen/proto/go_serv"
	"github.com/go-serv/foundation/internal/reflect"
	"github.com/go-serv/foundation/pkg/z"
)

var ref z.ReflectInterface

func init() {
	// Add the protobuf extensions.
	ref = reflect.NewReflection()
	ref.AddProtoExtension(go_serv.E_NewInsecureSession)
	ref.AddProtoExtension(go_serv.E_RequireSession)
	ref.AddProtoExtension(go_serv.E_OptionalSession)
	ref.AddProtoExtension(go_serv.E_CloseSession)
	ref.AddProtoExtension(go_serv.E_CopyMetaOff)
	ref.AddProtoExtension(go_serv.E_EncOff)
}
