package net

import (
	"context"
	net_call "github.com/go-serv/service/internal/grpc/context/net"
	net_req "github.com/go-serv/service/internal/grpc/request/net"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (mw *netMiddleware) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		var (
			md         metadata.MD
			ok         bool
			wrappedReq z.RequestInterface
			wrappedRes z.ResponseInterface
		)
		// Request metadata
		md, ok = metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "failed to retrieve request metadata")
		}
		//
		clientInfo := net_req.NewClientInfo(ctx)
		wrappedReq, err = net_req.NewRequest(req, md, clientInfo)
		if err != nil {
			return
		}
		//
		netCall := net_call.NewServerContext(ctx, wrappedReq, handler)
		wrappedRes, err = mw.newRequestChain().passThrough(netCall)
		if err != nil {
			return
		}
		//
		out, err = mw.newResponseChain().passThrough(wrappedRes)
		if err != nil {
			return
		}
		// Send response headers
		md, err = wrappedRes.Meta().Dehydrate()
		if err != nil {
			return
		}
		err = grpc.SendHeader(ctx, md)
		return
	}
}

func (mw *netMiddleware) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		return
	}
}
