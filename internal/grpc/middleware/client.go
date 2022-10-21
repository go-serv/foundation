package middleware

import (
	"context"
	"github.com/go-serv/foundation/internal/autogen/foundation"
	net_req "github.com/go-serv/foundation/internal/grpc/msg/request"
	net_res "github.com/go-serv/foundation/internal/grpc/msg/response"
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type clientMw struct {
	middleware
	client z.ClientInterface
}

func (m *clientMw) Client() z.ClientInterface {
	return m.client
}

func (m *clientMw) WithClient(client z.ClientInterface) {
	m.client = client
}

func (m *clientMw) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, in, out interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		var (
			req  z.RequestInterface
			res  z.ResponseInterface
			sref z.ServiceReflectInterface
		)

		client := m.Client()
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

		msg := in.(proto.Message)
		if sref, err = service.Reflection().ServiceReflectionFromMessage(msg); err != nil {
			return
		}

		// Pass server context through the request/response middleware chains.
		if err = m.requestPassThrough(netCtx, sref.FullName()); err != nil {
			return
		}
		if err = m.responsePassThrough(netCtx, sref.FullName()); err != nil {
			return
		}

		// Copy response metadata to the client if necessary.
		methodRef := req.MethodReflection()
		if methodRef.Has(foundation.E_ClientCopyMetaOff) {
			iv, _ := methodRef.Get(foundation.E_ClientCopyMetaOff)
			v := iv.(bool)
			if !v {
				res.Meta().Copy(m.Client().Meta())
			}
		} else {
			res.Meta().Copy(m.Client().Meta())
		}
		return
	}
}
