package z

import (
	"github.com/go-serv/foundation/pkg/z/platform"
	"google.golang.org/grpc"
)

type ServiceInterface interface {
	App() AppServerInterface
	BindApp(AppServerInterface)
	Endpoints() []EndpointInterface
	Name() string
	Codec() CodecInterface
	WithCodec(cc CodecInterface)
	Register(srv *grpc.Server)
}

type ServiceCfgInterface interface{}

type NetServiceCfgInterface interface {
	ServiceCfgInterface
}

type NetworkServiceInterface interface {
	ServiceInterface
	//EncriptionKey() []byte
	//WithEncriptionKey([]byte)
}

type LocalServiceInterface interface {
	ServiceInterface
}

type FtpUploadProfileInterface interface {
	RootDir() platform.Pathname
	MaxFileSize() int64
	FilePerms() platform.UnixPerms
}
