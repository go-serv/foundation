package pkg

import (
	"github.com/go-serv/service/internal/service"
	"google.golang.org/grpc"
	"log"
)

const (
	methodNotImplementedFmt = "method %s must be implemented"
)

type unimplementedMethod struct {
}

type baseService struct {
	service.BaseServiceInterface
	unimplementedMethod
}

type ServiceConfig interface {
	service.ConfigInterface
}

type LocalServiceInterface interface {
	service.LocalServiceInterface
}

func (s *baseService) CreateGrpcServer() *grpc.Server {
	log.Fatalf(methodNotImplementedFmt, "CreateGrpcServer")
	return nil
}
