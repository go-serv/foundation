package z

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/service/pkg/z/ancillary/crypto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ClientInterface interface {
	CodecAwareInterface
	MiddlewareAwareInterface
	ServiceName() protoreflect.FullName
	Endpoint() EndpointInterface
	ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	OnConnect(cc grpc.ClientConnInterface)
	WithDialOption(grpc.DialOption)
	DialOptions() []grpc.DialOption
	Meta() MetaInterface
	WithMeta(MetaInterface)
}

type NetworkClientInterface interface {
	ClientInterface
	NetService() NetworkServiceInterface
	BlockCipher() crypto.AEAD_CipherInterface
	WithBlockCipher(crypto.AEAD_CipherInterface)
}

type LocalClientInterface interface {
	ClientInterface
}
