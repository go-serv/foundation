package net

import (
	"github.com/go-serv/service/internal/ancillary/dictionary"
	"google.golang.org/grpc/metadata"
	"reflect"
	"strings"
)

type meta struct {
	data metadata.MD
	dic  interface{}
}

func (s *meta) Dictionary() interface{} {
	return s
}

type serverMeta struct {
	meta
}

func (s *meta) registerTypeHandlers(dic *HttpCommonDictionary) {
	dic.RegisterTypeHandler(reflect.TypeOf(""), func(op dictionary.OpType, name, alias string, rv reflect.Value) {
		switch op {
		case dictionary.HydrateOp:
			v := s.data.Get(name)
			if len(v) > 0 {
				rv.SetString(v[0])
			}
		case dictionary.DehydrateOp:
			v := rv.String()
			if len(v) > 0 {
				s.data.Set(name, v)
			}
		}
	})
}

func (s *serverMeta) Hydrate() error {
	dic := s.dic.(*HttpServerDictionary)
	return dic.Hydrate(dic)
}

type clientMeta struct {
	meta
}

func (c *clientMeta) Hydrate() error {
	dic := c.dic.(*HttpClientDictionary)
	return dic.Hydrate(dic)
}

func (m *serverMeta) GetIP() string {
	xForwardFor := m.data.Get("x-forwarded-for")
	if len(xForwardFor) > 0 && xForwardFor[0] != "" {
		ips := strings.Split(xForwardFor[0], ",")
		if len(ips) > 0 {
			clientIp := ips[0]
			return clientIp
		}
	}
	return ""
}
