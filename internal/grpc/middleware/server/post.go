// Post-request server middleware

package server

import (
	"context"
	"google.golang.org/grpc"
)

func PostUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		_, ok := grpc.Method(ctx)
		if !ok {
			return nil, nil // error
		}
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func PostStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		_, ok := grpc.MethodFromServerStream(ss)
		if !ok {
			return nil // error
		}
		return nil
	}
}
