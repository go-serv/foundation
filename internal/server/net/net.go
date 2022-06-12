package net

import "github.com/go-serv/service/pkg/z"

type netServer struct {
	z.ServerInterface
}

func (srv *netServer) Resolver() z.NetworkServerResolverInterface {
	return nil
}

func (srv *netServer) WithResolver(resolver z.NetworkServerResolverInterface) {

}
