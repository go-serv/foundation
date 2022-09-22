package middleware

import (
	"context"
	"github.com/go-serv/foundation/internal/autogen/go_serv/net/ext"
	net_req "github.com/go-serv/foundation/internal/grpc/msg/request"
	net_res "github.com/go-serv/foundation/internal/grpc/msg/response"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (chain *mwHandlersChain) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, in, out interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		var (
			req z.RequestInterface
			res z.ResponseInterface
		)

		client := chain.Client()
		if req, err = net_req.NewClientRequest(in, client.Meta(), nil); err != nil {
			return
		}

		md := metadata.MD{}
		opts = append(opts, grpc.Header(&md))
		if res, err = net_res.NewResponse(out, &md); err != nil {
			return
		}

		netCtx := ctx.(z.NetClientContextInterface)
		netCtx.WithClient(client.(z.NetworkClientInterface))
		netCtx.WithClientInvoker(invoker, cc, opts)
		netCtx.WithRequest(req)
		netCtx.WithResponse(res)

		// Pass server context through the request/response middleware chains.
		if err = chain.requestPassThrough(netCtx); err != nil {
			return
		}
		if err = chain.responsePassThrough(netCtx); err != nil {
			return
		}

		// Copy response metadata to the client if necessary.
		methodRef := req.MethodReflection()
		if methodRef.Has(ext.E_CopyMetaOff) {
			iv, _ := methodRef.Get(ext.E_CopyMetaOff)
			v := iv.(bool)
			if !v {
				res.Meta().Copy(chain.Client().Meta())
			}
		} else {
			res.Meta().Copy(chain.Client().Meta())
		}
		return
	}
}
