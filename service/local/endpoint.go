package local

import (
	"github.com/go-serv/foundation/internal/service"
	"net"
)

type localEndpoint struct {
	service.endpoint
}

func (ep *localEndpoint) listenAndServe() (err error) {
	var unixAddr *net.UnixAddr
	socketAddr := "@" + ep.Address()
	if unixAddr, err = net.ResolveUnixAddr("unix", socketAddr); err != nil {
		return
	}
	if ep.Lis, err = net.ListenUnix("unix", unixAddr); err != nil {
		return
	}
	return
}
