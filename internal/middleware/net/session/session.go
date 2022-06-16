package session

import (
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/grpc/meta/net"
	"github.com/go-serv/service/internal/grpc/session"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func serverSessionHandler(next z.NetChainElementFn, ctx z.NetContextInterface, req z.RequestInterface, res z.ResponseInterface) (err error) {
	if req.MethodReflection().Has(go_serv.E_RequireSession) {
		sId := req.Meta().Dictionary().(*net.HttpDictionary).SessionId
		sess := session.FindById(sId)
		if sess == nil || sess.State() != session.Active || sess.State() != session.New {
			return status.Error(codes.NotFound, "gRPC session not found or expired")
		}
		srvCtx := ctx.(z.NetServerContextInterface)
		srvCtx.WithSession(sess)
	}
	_, err = next(req, res)
	return
}

func clientSessionHandler(next z.NetChainElementFn, ctx z.NetContextInterface, req z.RequestInterface, res z.ResponseInterface) (err error) {
	if req.MethodReflection().Bool(go_serv.E_OptionalSession) {
		clntCtx := ctx.(z.NetClientContextInterface)
		client := clntCtx.Client()
		sId := client.Meta().Dictionary().(*net.HttpDictionary).SessionId
		req.Meta().Dictionary().(*net.HttpDictionary).SessionId = sId
	}
	_, err = next(req, res)
	return
}
