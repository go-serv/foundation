package server

import (
	"github.com/mesh-master/foundation/pkg/y/netparcel"
	"github.com/mesh-master/foundation/pkg/z"
	"github.com/mesh-master/foundation/pkg/z/dictionary"
	"github.com/mesh-master/foundation/pkg/z/platform"
	"github.com/mesh-master/foundation/service/net"
)

func NewNetParcel(eps []z.EndpointInterface, cfg ConfigInterface) *netParcel {
	svc := new(netParcel)
	svc.NetworkServiceInterface = net.NewNetworkService(Name, cfg, eps)
	svc.PostActions = make(map[string]netparcel.FtpPostActionHandlerFn)
	svc.RegisterFtpPostActionHandler(func(ctx dictionary.NetServerContextInterface, path platform.Pathname) error {
		return svc.handleGzipTarball(ctx, path)
	}, ".tar.gz")
	return svc
}
