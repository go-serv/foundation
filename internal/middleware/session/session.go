package session

import (
	"github.com/go-serv/foundation/internal/autogen/foundation"
	"github.com/go-serv/foundation/internal/grpc/session"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/dictionary"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errNotFound = status.Error(codes.NotFound, "gRPC session not found or expired")

func ServerRequestSessionHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	var (
		sess z.SessionInterface
	)
	srvCtx := ctx.(z.NetServerContextInterface)
	sId := req.Meta().Dictionary().(dictionary.BaseInterface).GetSessionId()
	if req.MethodReflection().Has(foundation.E_NewSession) { // Open a new session if necessary.
		v, _ := req.MethodReflection().Get(foundation.E_NewSession)
		lifetime := v.(uint32)
		if sId != 0 { // Assume session was opened by another call.
			sess = session.FindById(sId)
			if sess != nil {
				if ok := sess.IsValid(); !ok {
					return errNotFound
				}
			} else {
				return errNotFound
			}
		} else {
			sess = session.NewSession(lifetime)
		}
	} else if req.MethodReflection().Bool(foundation.E_RequireSession) { // Check if current session is valid.
		sess = session.FindById(sId)
		if sess == nil || !sess.IsValid() {
			return errNotFound
		}
	}
	if sess != nil {
		srvCtx.WithSession(sess)
	}
	_, err = next(req, nil)
	return
}

func ServerResponseSessionHandler(next z.NextHandlerFn, ctx z.NetContextInterface, res z.ResponseInterface) (err error) {
	if res.MethodReflection().Bool(foundation.E_CloseSession) { // Close current session.
		srvCtx := ctx.(z.NetServerContextInterface)
		if sess := srvCtx.Session(); sess != nil {
			sess.Close()
		}
	}
	_, err = next(nil, res)
	return
}

func ClientRequestSessionHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	if req.MethodReflection().Bool(foundation.E_RequireSession) {
		clntCtx := ctx.(z.NetClientContextInterface)
		client := clntCtx.Client()
		sId := client.Meta().Dictionary().(dictionary.BaseInterface).GetSessionId()
		req.Meta().Dictionary().(dictionary.BaseInterface).SetSessionId(sId)
	}
	_, err = next(req, nil)
	return
}
