package client

import (
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewClient(svcName protoreflect.FullName, e z.EndpointInterface) *client {
	c := new(client)
	c.svcName = svcName
	c.endpoint = e
	return c
}
