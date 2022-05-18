package runtime

import (
	"fmt"
	"github.com/go-serv/service/internal/service"
)

type netServiceRegistry map[string]service.NetworkServiceInterface

type runtime struct {
	registeredLocalService service.LocalServiceInterface
	registeredNetServices  netServiceRegistry
}

func (r *runtime) RegisterNetworkService(svcName string, svc service.NetworkServiceInterface) {
	if _, ok := r.registeredNetServices[svcName]; ok {
		panic(fmt.Sprintf("network service '%s' already registered", svcName))
	}
	r.registeredNetServices[svcName] = svc
}

func (r *runtime) RegisterLocalService(svcName string, svc service.LocalServiceInterface) {
	if r.registeredLocalService != nil {
		// @todo service name
		panic(fmt.Sprintf("Only one local service is allowed per application, '%s' already registered", ""))
	}
	r.registeredLocalService = svc
}

func (r *runtime) NetworkServices() []service.NetworkServiceInterface {
	var out []service.NetworkServiceInterface
	for _, svc := range r.registeredNetServices {
		out = append(out, svc)
	}
	return out
}

func Runtime() *runtime {
	return rt
}
