// Pre-request server middleware
// Creates a descriptor for a gRPC call holding meta information about the call

package server_md

import (
	"context"
	"google.golang.org/grpc"
)

func PreUnaryInterceptor(svcName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func PreStreamInterceptor(svcName string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return nil
	}
}
