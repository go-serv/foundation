package server

import (
	"github.com/mesh-master/foundation/pkg/z"
)

func NewWebproxy(ep z.NetEndpointInterface) z.WebProxyInterface {
	wp := new(webproxy)
	wp.endpoint = ep
	return wp
}
