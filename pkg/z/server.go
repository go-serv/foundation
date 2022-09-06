package z

import (
	job "github.com/AgentCoop/go-work"
	"google.golang.org/grpc"
)

type EndpointInterface interface {
	Address() string
	ServeTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	GrpcServer() *grpc.Server
	BindGrpcServer(*grpc.Server)
	GrpcServerOptions() []grpc.ServerOption
	WithGrpcServerOptions(opts ...grpc.ServerOption)
	Service() ServiceInterface
	BindService(ServiceInterface)
}

type ServerInterface interface {
	CodecAwareInterface
	AddGrpcServerOption(opt grpc.ServerOption)
	GrpcServerOptions() []grpc.ServerOption
	AddEndpoint(endpoint EndpointInterface)
	Endpoints() []EndpointInterface
	Start()
	Stop()
	MainJob() job.JobInterface
	MiddlewareGroup() MiddlewareInterface
	WithMiddlewareGroup(mg MiddlewareInterface)
}

type LocalServerInterface interface {
	ServerInterface
}

type NetworkServerInterface interface {
	ServerInterface
	Resolver() NetworkServerResolverInterface
	WithResolver(NetworkServerResolverInterface)
}

type ServerResolverInterface interface {
}

type AccessTokenVerifierFn func(AccessTokenInterface) bool

type NetworkServerResolverInterface interface {
	ServerResolverInterface
	VerifyAccessToken(fn AccessTokenVerifierFn) (bool, error)
}
