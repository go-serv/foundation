package y

import (
	net_mw "github.com/go-serv/service/internal/grpc/mw_group/net"
	"github.com/go-serv/service/pkg/z"
)

func NewMiddleware() z.NetMiddlewareInterface {
	return net_mw.NewMiddleware()
}
