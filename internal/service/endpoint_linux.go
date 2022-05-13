package service

import "net"

func NewLocalEndpoint(s LocalServiceInterface, pathname string) *localEndpoint {
	pathname = "@" + pathname // Listen on an abstract unix domain socket
	e := &localEndpoint{NewEndpoint(s), pathname}
	return e
}

func (e *localEndpoint) Listen() error {
	var err error
	var unixAddr *net.UnixAddr
	unixAddr, err = net.ResolveUnixAddr(udsNetwork, e.pathname)
	if err != nil {
		return err
	}
	e.lis, err = net.ListenUnix(udsNetwork, unixAddr)
	if err != nil {
		return err
	}
	return nil
}
