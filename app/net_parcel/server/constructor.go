package server

import (
	"github.com/go-serv/foundation/pkg/y/netparcel"
	"github.com/go-serv/foundation/pkg/z"
	"github.com/go-serv/foundation/pkg/z/dictionary"
	"github.com/go-serv/foundation/pkg/z/platform"
	"github.com/go-serv/foundation/service/net"
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
