package net

import "github.com/go-serv/foundation/pkg/z"

type cfg struct {
	z.NetServiceCfgInterface
	*WebProxyConfig
}
