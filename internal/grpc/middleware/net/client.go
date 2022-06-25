package net

import (
	"context"
	"github.com/go-serv/service/internal/autogen/proto/go_serv"
	"github.com/go-serv/service/internal/grpc/codec"
	net_req "github.com/go-serv/service/internal/grpc/msg/request"
	net_res "github.com/go-serv/service/internal/grpc/msg/response"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

func (mw *netMiddleware) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, reqp, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		var (
			req          z.RequestInterface
			res          z.ResponseInterface
			dfReq, dfRes z.DataFrameInterface
		)
		if dfReq, err = codec.NewDataFrame(reqp); err != nil {
			return
		}
		if dfRes, err = codec.NewDataFrame(resp); err != nil {
			return
		}
		clnt := mw.Client()
		if req, err = net_req.NewClientRequest(dfReq, clnt.Meta(), nil); err != nil {
			return
		}
		md := metadata.MD{}
		opts = append(opts, grpc.Header(&md))
		if res, err = net_res.NewResponse(dfRes, &md); err != nil {
			return
		}
		netCtx := ctx.(z.NetClientContextInterface)
		netCtx.WithClient(clnt.(z.NetworkClientInterface))
		netCtx.WithInput(reqp.(proto.Message))
		netCtx.WithOutput(resp.(proto.Message))
		netCtx.WithClientInvoker(invoker, cc, opts)
		netCtx.WithRequest(req)
		netCtx.WithResponse(res)
		// Request chain.
		//reqp.(proto.Message).ProtoReflect().
		//reflect.ValueOf(reqp).Elem().Set(reflect.ValueOf(dfReq.ProtoMessage()).Elem())
		if err = mw.newRequestChain().passThrough(netCtx); err != nil {
			return
		}
		// Response chain.
		err = mw.newResponseChain().passThrough(netCtx)
		// Copy response metadata to the client if necessary.
		methodRef := req.MethodReflection()
		if methodRef.Has(go_serv.E_CopyMetaOff) {
			iv, _ := methodRef.Get(go_serv.E_CopyMetaOff)
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
