package runtime

import (
	i "github.com/go-serv/service/internal"
)

type RuntimeInterface interface {
	NetworkServices() []i.NetworkServiceInterface
}
