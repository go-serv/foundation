package runtime

import (
	"fmt"
	"github.com/go-serv/service/internal/service/local"
	"github.com/go-serv/service/internal/service/net"
)

type netServiceRegistry map[string]net.NetworkServiceInterface

type runtime struct {
	registeredLocalService local.LocalServiceInterface
	registeredNetServices  netServiceRegistry
}

func (r *runtime) RegisterNetworkService(svcName string, svc net.NetworkServiceInterface) {
	if _, ok := r.registeredNetServices[svcName]; ok {
		panic(fmt.Sprintf("network service '%s' already registered", svcName))
	}
	r.registeredNetServices[svcName] = svc
}

func (r *runtime) RegisterLocalService(svcName string, svc local.LocalServiceInterface) {
	if r.registeredLocalService != nil {
		// @todo service name
		panic(fmt.Sprintf("Only one local service is allowed per application, '%s' already registered", ""))
	}
	r.registeredLocalService = svc
}

func (r *runtime) NetworkServices() []net.NetworkServiceInterface {
	var out []net.NetworkServiceInterface
	for _, svc := range r.registeredNetServices {
		out = append(out, svc)
	}
	return out
}

func Runtime() *runtime {
	return rt
}
