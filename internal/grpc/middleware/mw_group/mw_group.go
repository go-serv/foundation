package mw_group

import (
	"context"
	"github.com/go-serv/service/internal/grpc/request"
	"github.com/go-serv/service/internal/grpc/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type RequestMiddlewareHandler func(r request.RequestInterface, svc interface{}) error
type ResponseMiddlewareHandler func(r response.ResponseInterface, svc interface{}) error

type GroupItem struct {
	reqHandler RequestMiddlewareHandler
	resHandler ResponseMiddlewareHandler
}

type mwGroup struct {
	items []*GroupItem
}

func (mw *mwGroup) AddItem(item *GroupItem) {
	mw.items = append(mw.items, item)
}

func (mw *mwGroup) NetUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var md metadata.MD
		var ok bool
		var ctxPeer *peer.Peer
		// Request metadata
		md, ok = metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Internal, "failed to retrieve metadata")
		}
		// Client address
		ctxPeer, ok = peer.FromContext(ctx)
		if !ok {
			return nil, nil
		}
		_ = ctxPeer
		r := request.NewNetUnaryRequest(ctx, req, md)
		for _, item := range mw.items {
			err := item.reqHandler(r, info.Server)
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
		for i := len(mw.items) - 1; i >= 0; i-- {
			err := mw.items[i].resHandler(res, info.Server)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	}
}
