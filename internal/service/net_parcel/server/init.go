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

func (svc *netParcel) processTarball(path platform.Pathname) (err error) {
	var untar z.RunnableInterface
	if untar, err = archive.NewUntar(path, ancillary.GzipCompressor, svc.ArchiveOptions); err != nil {
		return
	}
	err = untar.Run()
	return
}

func init() {
	svc := new(netParcel)
	svc.NetworkServiceInterface = net.NewNetworkService(Name)
	svc.PostActions = make(map[string]netparcel.FtpPostActionHandlerFn)
	svc.RegisterFtpPostActionHandler(func(path platform.Pathname) error {
		return nil //svc.processTarball(path)
	}, ".tar.gz")
	rt := runtime.Runtime()
	rt.RegisterNetworkService(svc)
	rt.Reflection().AddService(Name)
	rt.Reflection().Populate()
}
