package net

import (
	"context"
	z "github.com/go-serv/service/internal"
	net_call "github.com/go-serv/service/internal/grpc/context/net"
	net_req "github.com/go-serv/service/internal/grpc/request/net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (mw *netMwGroup) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		var (
			md         metadata.MD
			ok         bool
			wrappedReq z.RequestInterface
			res        z.ResponseInterface
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
		netCall := net_call.NewCall(ctx, wrappedReq, handler)
		res, err = mw.newRequestChain().passThrough(netCall)
		if err != nil {
			return
		}
		//
		out, err = mw.newResponseChain().passThrough(res)
		return
	}
}

func (mw *netMwGroup) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		return
	}
}

func (mw *netMwGroup) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return nil
	}
}
