package cerberus

import (
	"github.com/go-serv/foundation/internal/autogen/net/cerberus"
	"github.com/go-serv/foundation/internal/service"
)

func init() {
	service.Reflection().AddProtoExtension(cerberus.E_RequiresRole)
}
