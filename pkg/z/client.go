package z

import (
	job "github.com/AgentCoop/go-work"
	"github.com/go-serv/foundation/pkg/z/ancillary/crypto"
	"google.golang.org/grpc"
)

type ClientInterface interface {
	CodecAwareInterface
	MiddlewareAwareInterface
	ServiceName() string
	Endpoint() EndpointInterface
	ConnectTask(j job.JobInterface) (job.Init, job.Run, job.Finalize)
	OnConnect(cc grpc.ClientConnInterface)
	WithDialOption(grpc.DialOption)
	DialOptions() []grpc.DialOption
	Meta() MetaInterface
	WithMeta(MetaInterface)
	// WithOptions sets options for a call.
	WithOptions(any)
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
