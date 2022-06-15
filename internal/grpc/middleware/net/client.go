package net

import (
	"context"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	net_req "github.com/go-serv/service/internal/grpc/request/net"
	net_res "github.com/go-serv/service/internal/grpc/response/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (mw *netMiddleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		var (
			wrappedReq z.RequestInterface
			wrappedRes z.ResponseInterface
		)
		clnt := mw.Client()
		wrappedReq, err = net_req.NewClientRequest(req, clnt.Meta(), nil)
		if err != nil {
			return
		}
		md := metadata.MD{}
		opts = append(opts, grpc.Header(&md))
		wrappedRes = net_res.NewResponse(reply, &md)
		netCtx := ctx.(z.NetClientContextInterface)
		dd := clnt.(z.NetworkClientInterface)
		_ = dd
		netCtx.WithClient(clnt.(z.NetworkClientInterface))
		netCtx.WithClientInvoker(invoker, cc, opts)
		netCtx.WithRequest(wrappedReq)
		netCtx.WithResponse(wrappedRes)
		// Request chain
		_, err = mw.newRequestChain().passThrough(netCtx)
		if err != nil {
			return
		}
		// Response chain
		_, err = mw.newResponseChain().passThrough(wrappedRes)
		// Copy response metadata to the client if necessary.
		mref := wrappedReq.MethodReflection()
		if mref.Has(go_serv.E_CopyMetaOff) {
			iv, _ := mref.Get(go_serv.E_CopyMetaOff)
			v := iv.(bool)
			if !v {
				wrappedRes.Meta().Copy(mw.Client().Meta())
			}
		} else {
			wrappedRes.Meta().Copy(mw.Client().Meta())
		}
		return
	}
}
