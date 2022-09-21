package net

import (
	"context"
	ext "github.com/go-serv/foundation/internal/autogen/go_serv/net/ext"
	net_req "github.com/go-serv/foundation/internal/grpc/msg/request"
	net_res "github.com/go-serv/foundation/internal/grpc/msg/response"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (mw *netMiddleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, ifreq, ifreplay interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		var (
			req z.RequestInterface
			res z.ResponseInterface
		)
		client := mw.Client()
		if req, err = net_req.NewClientRequest(ifreq.(z.DataFrameInterface), client.Meta(), nil); err != nil {
			return
		}
		md := metadata.MD{}
		opts = append(opts, grpc.Header(&md))
		if res, err = net_res.NewResponse(ifreplay.(z.DataFrameInterface), &md); err != nil {
			return
		}
		netCtx := ctx.(z.NetClientContextInterface)
		netCtx.WithClient(client.(z.NetworkClientInterface))
		netCtx.WithClientInvoker(invoker, cc, opts)
		netCtx.WithRequest(req)
		netCtx.WithResponse(res)
		// Request chain.
		if err = mw.newRequestChain().passThrough(netCtx); err != nil {
			return
		}
		// Response chain.
		err = mw.newResponseChain().passThrough(netCtx)
		// Copy response metadata to the client if necessary.
		methodRef := req.MethodReflection()
		if methodRef.Has(ext.E_CopyMetaOff) {
			iv, _ := methodRef.Get(ext.E_CopyMetaOff)
			v := iv.(bool)
			if !v {
				res.Meta().Copy(mw.Client().Meta())
			}
		} else {
			res.Meta().Copy(mw.Client().Meta())
		}
		return
	}
}
