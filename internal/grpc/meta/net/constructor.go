package net

import (
	"github.com/mesh-master/foundation/pkg/z/dictionary"
	"google.golang.org/grpc/metadata"
)

func NewRequestMeta(data *metadata.MD) *requestMeta {
	m := new(requestMeta)
	m.data = data
	m.DictionaryInterface = dictionary.NewNetRequestDictionary()
	return m
}

func NewResponseMeta(data *metadata.MD) *responseMeta {
	m := new(responseMeta)
	m.data = data
	m.DictionaryInterface = dictionary.NewNetResponseDictionary()
	return m
}
