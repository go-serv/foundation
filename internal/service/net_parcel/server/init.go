package server

import (
	"github.com/go-serv/service/internal/ancillary/archive"
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/internal/service/net"
	"github.com/go-serv/service/pkg/y/netparcel"
	"github.com/go-serv/service/pkg/z"
	"github.com/go-serv/service/pkg/z/ancillary"
	"github.com/go-serv/service/pkg/z/platform"
)

func (svc *netParcel) handleGzipTarball(ctx z.NetServerContextInterface, path platform.Pathname) (err error) {
	var (
		plat  z.PlatformInterface
		untar z.RunnableInterface
	)
	if ctx.Tenant() != 0 {
		// TODO: retrieve a tenant platform API object from the runtime registry
	} else {
		plat = runtime.Runtime().Platform()
	}
	if untar, err = archive.NewUntar(plat, path, ancillary.GzipCompressor, svc.ArchiveOptions); err != nil {
		return
	}
	if err = untar.Run(); err != nil {
		return
	}
	// Probably there won't be a case that will require to keep the uploaded tarball once it's unpacked.
	// We can remove it.
	err = plat.Remove(path)
	return
}

func init() {
	svc := new(netParcel)
	svc.NetworkServiceInterface = net.NewNetworkService(Name)
	svc.PostActions = make(map[string]netparcel.FtpPostActionHandlerFn)
	svc.RegisterFtpPostActionHandler(func(ctx z.NetServerContextInterface, path platform.Pathname) error {
		return svc.handleGzipTarball(ctx, path)
	}, ".tar.gz")
	rt := runtime.Runtime()
	rt.RegisterNetworkService(svc)
	rt.Reflection().AddService(Name)
	rt.Reflection().Populate()
}
