package net

import "google.golang.org/grpc/metadata"

func NewMeta(data *metadata.MD) *meta {
	m := new(meta)
	m.data = data
	m.dic = new(HttpDictionary)
	m.registerTypeHandlers(m.dic)
	return m
}
