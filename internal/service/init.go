package service

import (
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/autogen/go_serv/net/ext"
	"github.com/go-serv/foundation/internal/reflect"
	"github.com/go-serv/foundation/pkg/z"
)

var ref z.ReflectInterface

func init() {
	// Add the protobuf extensions.
	ref = reflect.NewReflection()
	ref.AddProtoExtension(ext.E_NewInsecureSession)
	ref.AddProtoExtension(ext.E_RequireSession)
	ref.AddProtoExtension(ext.E_OptionalSession)
	ref.AddProtoExtension(ext.E_CloseSession)
	ref.AddProtoExtension(ext.E_CopyMetaOff)
	ref.AddProtoExtension(ext.E_EncOff)
	ref.AddProtoExtension(foundation.E_RequireSession)
}
