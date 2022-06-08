package net

import "strings"

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
