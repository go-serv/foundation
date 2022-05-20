package runtime

import (
	"fmt"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/service/local"
)

type netServiceRegistry map[string]i.NetworkServiceInterface

type runtime struct {
	registeredLocalService local.LocalServiceInterface
	registeredNetServices  netServiceRegistry
}

func (r *runtime) RegisterNetworkService(svcName string, svc i.NetworkServiceInterface) {
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

func (r *runtime) NetworkServices() []i.NetworkServiceInterface {
	var out []i.NetworkServiceInterface
	for _, svc := range r.registeredNetServices {
		out = append(out, svc)
	}
	return out
}

func Runtime() *runtime {
	return rt
}
