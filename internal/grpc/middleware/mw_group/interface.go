package mw_group

import "google.golang.org/grpc"

type MiddlewareGroupInterface interface {
	NetUnaryInterceptor() grpc.UnaryServerInterceptor
}
