package runtime

import "github.com/go-serv/service/internal/service"

type RuntimeInterface interface {
	NetworkServices() []service.NetworkServiceInterface
}
