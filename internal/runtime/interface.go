package runtime

import "github.com/go-serv/service/internal/service/net"

type RuntimeInterface interface {
	NetworkServices() []net.NetworkServiceInterface
}
