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
	if req.ServiceReflection().Bool(foundation.E_EnforceSecureChannel) {
		if !srvCtx.NetworkService().IsTlsEnabled() {
			sess = srvCtx.Session()
			if sess == nil || sess.BlockCipher() == nil {
				errMsg := fmt.Sprintf("gRPC '%s' call requires a secure channel", req.MethodReflection().Descriptor().FullName())
				err = status.Error(codes.FailedPrecondition, errMsg)
				return
			}
		}
	}

	if !verifyApiKey(srvCtx, req) {
		err = status.Error(codes.Unauthenticated, "invalid API key")
		return
	}

	_, err = next(req, nil)
	return
}

func ServerResponseNetHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	return
}

func ClientRequestNetHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	if v, has := req.ServiceReflection().Get(foundation.E_AuthType); has {
		authType := v.(foundation.AuthType)
		if authType == foundation.AuthType_ApiKey {
			apiKey := ctx.(z.NetworkClientInterface).ApiKey()
			if len(apiKey) == 0 {
				panic("empty api key")
			}
		}
	}
	_, err = next(req, nil)
	return
}

func ClientResponseNetHandler(next z.NextHandlerFn, _ z.NetContextInterface, res z.ResponseInterface) (err error) {
	_, err = next(nil, res)
	return
}
