package net

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

func (mw *netMiddleware) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, ifreq interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
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
		if req, err = request.NewServerRequest(ifreq.(z.DataFrameInterface), &md, clientInfo); err != nil {
			return
		}
		//
		srvCxt := net_call.NewServerContext(ctx, req, handler)
		if err = mw.newRequestChain().passThrough(srvCxt); err != nil {
			return
		}
		if err = mw.newResponseChain().passThrough(srvCxt); err != nil {
			return
		}
		// Send response headers
		if md, err = srvCxt.Response().Meta().Dehydrate(); err != nil {
			return
		}
		if err = grpc.SendHeader(ctx, md); err != nil {
			return
		}
		//
		out = srvCxt.Response().DataFrame()
		return
	}
}

func (mw *netMiddleware) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		return
	}
}
