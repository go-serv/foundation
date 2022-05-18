package mw_group

import (
	"context"
	"github.com/go-serv/service/internal/grpc/request"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type MiddlewareHandler func(r request.RequestInterface) error

type mwGroup struct {
	items []MiddlewareHandler
}

func (mw *mwGroup) AddItem(handler MiddlewareHandler) {
	mw.items = append(mw.items, handler)
}

func (mw *mwGroup) NetUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, nil
		}
		r := request.NewNetUnaryRequest(ctx, req, md)
		for _, mwFn := range mw.items {
			err := mwFn(r)
			if err != nil {
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}
