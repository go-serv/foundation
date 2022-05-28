package net

import (
	"context"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/grpc/mw_group"
	req_net "github.com/go-serv/service/internal/grpc/request/net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type netMwGroup struct {
	mw_group.MwGroup
}

func (mw *netMwGroup) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//var md metadata.MD
		var ok bool
		var ctxPeer *peer.Peer
		// Request metadata
		//md, ok = metadata.FromIncomingContext(ctx)
		//if !ok {
		//	return nil, status.Error(codes.Internal, "failed to retrieve metadata")
		//}
		// Client address
		ctxPeer, ok = peer.FromContext(ctx)
		if !ok {
			return nil, nil
		}
		_ = ctxPeer
		r := req_net.FromServerContext(ctx, req)
		for _, item := range mw.Items {
			err := item.ReqHandler(r, info.Server)
			if err != nil {
				return nil, err
			}
		}
		// Handle request
		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		// Iterate over the response middleware handlers in reverse order
		for i := len(mw.Items) - 1; i >= 0; i-- {
			err := mw.Items[i].ResHandler(res, info.Server)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	}
}

func (mw *netMwGroup) UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		r := req_net.FromClientContext(ctx, req, method)
		svc := mw.Target.(i.NetworkClientInterface).NetService()
		// Invoke middleware request handlers
		for _, item := range mw.Items {
			err := item.ReqHandler(r, svc)
			if err != nil {
				return nil
			}
		}
		// Invoke gRPC method
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			return err
		}
		// Invoke middleware response handlers in reverse order
		for i := len(mw.Items) - 1; i >= 0; i-- {
			err := mw.Items[i].ResHandler(reply, svc)
			if err != nil {
				return nil
			}
		}
		return nil
	}
}
