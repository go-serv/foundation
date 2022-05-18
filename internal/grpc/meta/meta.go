package meta

import (
	"google.golang.org/grpc/metadata"
	"strings"
)

type meta struct {
	data        metadata.MD
	cryptoAlgo  string
	cryptoNonce string
}

type serverMeta struct {
	meta
}

type clientMeta struct {
	meta
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
