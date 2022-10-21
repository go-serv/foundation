package middleware

import (
	"github.com/go-serv/foundation/internal/autogen/net/cerberus"
	"github.com/go-serv/foundation/pkg/z"
)

func ServerReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	//srvCtx := ctx.(z.NetServerContextInterface
	if req.ServiceReflection().Has(cerberus.E_RequiresRole) {
		roles, _ := req.ServiceReflection().Get(cerberus.E_RequiresRole)
		_ = roles
	}
	_, err = next(req, nil)
	return
}
