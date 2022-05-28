package server

import (
	i "github.com/go-serv/service/internal"
)

func NewLocalEndpoint(svc i.LocalServiceInterface) *localEndpoint {
	pathname := string(svc.Name())
	e := &localEndpoint{NewEndpoint(), pathname}
	return e
}
