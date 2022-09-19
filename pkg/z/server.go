package z

import (
	"crypto/tls"
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/pkg/z/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

type NetEndpointInterface interface {
	EndpointInterface
	IsSecure() bool
	TlsConfig() *tls.Config
	TransportCredentials() credentials.TransportCredentials
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

type AccessTokenVerifierFn func(security.AccessTokenInterface) bool

type NetworkServerResolverInterface interface {
	ServerResolverInterface
	VerifyAccessToken(fn AccessTokenVerifierFn) (bool, error)
}
