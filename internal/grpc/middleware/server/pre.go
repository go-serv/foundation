// Pre-request server middleware
// Creates a descriptor for a gRPC call holding meta information about the call

package server

import (
	"context"
	"google.golang.org/grpc"
)

//func handle(ctx string, header metadata.MD) (calldesc.ServerCallDescriptor, error) {
//	desc := calldesc.NewServerCallDesc()
//	return desc, nil
//}

func PreUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//var (
		//	methodName string
		//	header     metadata.MD
		//	ok         bool
		//)
		//methodName, ok = grpc.Method(ctx)
		//desc, err := handle(methodName, header)
		//if err != nil {
		//	return nil, err
		//}
		return handler(ctx, req)
	}
}

func PreStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return nil
	}
}
