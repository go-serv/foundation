package client

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewClient(svcName protoreflect.FullName, e i.EndpointInterface) *client {
	c := new(client)
	c.svcName = svcName
	c.endpoint = e
	return c
}
