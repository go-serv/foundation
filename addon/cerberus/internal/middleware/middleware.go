package middleware

import (
	"github.com/mesh-master/foundation/internal/autogen/foundation"
	"github.com/mesh-master/foundation/pkg/z"
)

func ServerReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	//srvCtx := ctx.(z.NetServerContextInterface
	if req.ServiceReflection().Has(foundation.E_RequiresRole) {
		roles, _ := req.ServiceReflection().Get(foundation.E_RequiresRole)
		_ = roles
	}
	_, err = next(req, nil)
	return
}
