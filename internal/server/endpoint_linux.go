package server

//func NewLocalEndpoint(s service.LocalServiceInterface, pathname string) *Server.Server.localEndpoint {
//	pathname = "@" + pathname // Listen on an abstract unix domain socket
//	e := &Server.localEndpoint{service.NewEndpoint(s), pathname}
//	return e
//}
//
//func (e *Server.Server.localEndpoint) Listen() error {
//	var err error
//	var unixAddr *net.UnixAddr
//	unixAddr, err = net.ResolveUnixAddr(udsNetwork, e.pathname)
//	if err != nil {
//		return err
//	}
//	e.lis, err = net.ListenUnix(udsNetwork, unixAddr)
//	if err != nil {
//		return err
//	}
//	return nil
//}
