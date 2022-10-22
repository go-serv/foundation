package net

import (
	"fmt"
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServerRequestNetHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		sess z.SessionInterface
	)
	srvCtx := ctx.(z.NetServerContextInterface)
	if v, has := req.ServiceReflection().Shadow(foundation.E_EnforceSecureChannel, foundation.E_MEnforceSecureChannel, req.MethodReflection()); has {
		enforce := v.(bool)
		if enforce && !srvCtx.NetworkService().IsTlsEnabled() {
			sess = srvCtx.Session()
			if sess == nil || sess.BlockCipher() == nil {
				errMsg := fmt.Sprintf("gRPC '%s' call requires a secure channel", req.MethodReflection().Descriptor().FullName())
				err = status.Error(codes.FailedPrecondition, errMsg)
				return
			}
		}
	}
	_, err = next(req, nil)
	return
}

func ServerResponseNetHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	return
}

func ClientRequestNetHandler(next z.NextHandlerFn, _ z.NetContextInterface, req z.RequestInterface) (err error) {
	_, err = next(req, nil)
	return
}

func ClientResponseNetHandler(next z.NextHandlerFn, _ z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	return
}
