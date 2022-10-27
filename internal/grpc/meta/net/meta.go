package net

import (
	"github.com/go-serv/foundation/pkg/ancillary/struc/dictionary"
	"github.com/go-serv/foundation/pkg/z"
	dic_defs "github.com/go-serv/foundation/pkg/z/dictionary"
	"google.golang.org/grpc/metadata"
)

type meta struct {
	dictionary.DictionaryInterface
	data *metadata.MD
}

type requestMeta struct {
	meta
}

type responseMeta struct {
	meta
}

func (m *meta) Dictionary() dictionary.DictionaryInterface {
	return m.DictionaryInterface
}

func (m *meta) WithDictionary(d dictionary.DictionaryInterface) {
	m.DictionaryInterface = d
}

func (m *meta) Get(k string) (v string, has bool) {
	values := m.data.Get(k)
	if len(values) > 0 {
		return values[0], true
	} else {
		return "", false
	}
}

func (m *meta) Set(k string, v string) {
	m.data.Set(k, v)
}

func (m *meta) Copy(dst z.MetaInterface) {
	src := m.Dictionary().(dic_defs.BaseInterface)
	dst.Dictionary().(dic_defs.BaseInterface).SetSessionId(src.GetSessionId())
}

func (m *meta) Hydrate() error {
	return dictionary.Dictionary{}.Import(m)
}

func (m *meta) Dehydrate() (md metadata.MD, err error) {
	if m.data == nil {
		m.data = &metadata.MD{}
	}
	err = dictionary.Dictionary{}.Export(m)
	if err != nil {
		return nil, err
	}
	return *m.data, nil
}
