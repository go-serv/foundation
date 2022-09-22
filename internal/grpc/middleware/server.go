package middleware

import (
	"context"
	net_call "github.com/go-serv/foundation/internal/grpc/callctx/net"
	"github.com/go-serv/foundation/internal/grpc/msg/request"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (chain *mwHandlersChain) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, in interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		var (
			md  metadata.MD
			ok  bool
			req z.RequestInterface
		)
		if md, ok = metadata.FromIncomingContext(ctx); !ok {
			return nil, status.Error(codes.Internal, "failed to retrieve request metadata")
		}

		//
		clientInfo := request.NewClientInfo(ctx)
		if req, err = request.NewServerRequest(in, &md, clientInfo); err != nil {
			return
		}

		// Pass server context through the request/response middleware chains.
		srvCxt := net_call.NewServerContext(ctx, req, handler)
		if err = chain.requestPassThrough(srvCxt); err != nil {
			return
		}
		if err = chain.responsePassThrough(srvCxt); err != nil {
			return
		}

		// Send response headers.
		if md, err = srvCxt.Response().Meta().Dehydrate(); err != nil {
			return
		}
		if err = grpc.SendHeader(ctx, md); err != nil {
			return
		}
		out = srvCxt.Response().Data()
		return
	}
}

func (chain *mwHandlersChain) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		return
	}
}
