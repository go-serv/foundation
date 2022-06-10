package net

import "google.golang.org/grpc/metadata"

func NewMeta(md metadata.MD) *meta {
	m := new(meta)
	//if md == nil {
	//	m.data = metadata.MD{}
	//} else {
	m.data = md
	//}
	m.dic = new(HttpDictionary)
	m.registerTypeHandlers(m.dic)
	return m
}
