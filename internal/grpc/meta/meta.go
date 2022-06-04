package meta

import (
	"google.golang.org/grpc/metadata"
	"strings"
)

type meta struct {
	data metadata.MD
}

type serverMeta struct {
	meta
	srvDic *HttpServerDictionary
}

func (s *serverMeta) Hydrate() error {
	return s.srvDic.Hydrate(s.srvDic)
}

type clientMeta struct {
	meta
	clientDic *HttpClientDictionary
}

func (c *clientMeta) Hydrate() error {
	return c.clientDic.Hydrate(c.clientDic)
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
