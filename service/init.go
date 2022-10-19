package service

import (
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/service"
)

func init() {
	service.Reflection().AddProtoExtension(foundation.E_RequireSession)
}
