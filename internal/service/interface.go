package service

import (
	"google.golang.org/grpc"
)

type ConfigInterface interface {
}

type GrpcMessageWrapperFn func(in []byte) []byte

type baseServiceInterface interface {
	Service_Register(srv *grpc.Server)
}

type LocalServiceInterface interface {
	baseServiceInterface
}

type NetworkServiceInterface interface {
	baseServiceInterface
}
