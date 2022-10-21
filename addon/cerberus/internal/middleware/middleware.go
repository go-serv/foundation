package middleware

import (
	"github.com/go-serv/foundation/pkg/z"
)

func ServerReqHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	//srvCtx := ctx.(z.NetServerContextInterface)
	_, err = next(req, nil)
	return
}
