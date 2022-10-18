package session

import (
	"github.com/go-serv/foundation/internal/autogen/go_serv/net/ext"
	"github.com/go-serv/foundation/internal/grpc/meta/net"
	"github.com/go-serv/foundation/internal/grpc/session"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ServerRequestSessionHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	srvCtx := ctx.(z.NetServerContextInterface)
	sId := req.Meta().Dictionary().(*net.HttpDictionary).SessionId
	if req.MethodReflection().Bool(ext.E_RequireSession) {
		sess := session.FindById(sId)
		if sess == nil || (sess.State() != session.Active && sess.State() != session.New) {
			return status.Error(codes.NotFound, "gRPC session not found or expired")
		}
		srvCtx.WithSession(sess)
	} else if req.MethodReflection().Bool(ext.E_OptionalSession) {
		sess := session.FindById(sId)
		srvCtx.WithSession(sess)
	}
	_, err = next(req, nil)
	return
}

func ClientRequestSessionHandler(next z.NextHandlerFn, ctx z.NetContextInterface, req z.RequestInterface) (err error) {
	if req.MethodReflection().Bool(ext.E_OptionalSession) {
		clntCtx := ctx.(z.NetClientContextInterface)
		client := clntCtx.Client()
		sId := client.Meta().Dictionary().(*net.HttpDictionary).SessionId
		req.Meta().Dictionary().(*net.HttpDictionary).SessionId = sId
	}
	_, err = next(req, nil)
	return
}
