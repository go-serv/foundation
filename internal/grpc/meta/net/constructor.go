package net

import "google.golang.org/grpc/metadata"

func NewServerMeta(md metadata.MD) *serverMeta {
	s := new(serverMeta)
	s.data = md
	s.dic = new(HttpServerDictionary)
	s.registerTypeHandlers()
	return s
}

func NewClientMeta() *clientMeta {
	c := new(clientMeta)
	c.data = metadata.MD{}
	c.dic = new(HttpClientDictionary)
	return c
}
