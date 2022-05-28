package server

import (
	i "github.com/go-serv/service/internal"
)

func NewLocalEndpoint(svc i.LocalServiceInterface) *localEndpoint {
	pathname := "@" + string(svc.Name()) // Listen on an abstract unix domain socket
	e := &localEndpoint{NewEndpoint(), pathname}
	return e
}
