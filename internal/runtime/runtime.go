package runtime

import (
	"errors"
	"fmt"
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/service/local"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMethodDescriptorNotFound = errors.New("")
)

type (
	netServiceRegistry   map[string]i.NetworkServiceInterface
	localServiceRegistry map[string]i.LocalServiceInterface
)

type runtime struct {
	registeredLocalService localServiceRegistry
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
	//r.registeredLocalService = svc
}

func (r *runtime) allRegisteredServices() []interface{} {
	all := make([]interface{}, len(r.registeredLocalService)+len(r.registeredLocalService))
	for _, v := range r.registeredNetServices {
		all = append(all, v)
	}
	for _, v := range r.registeredLocalService {
		all = append(all, v)
	}
	return all
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

func (r *runtime) MethodDescriptorByMessage(msg proto.Message) (i.MethodDescriptorInterface, error) {
	key := msg.ProtoReflect().Descriptor().FullName()
	//
	for _, svc := range r.allRegisteredServices() {
		svcDesc := svc.(i.BaseServiceInterface).Service_Descriptor()
		methods := svcDesc.Descriptor().Methods()
		l1 := methods.Len()
		for ii := 0; ii < l1; ii++ {
			m := methods.Get(ii)
			input := m.Input().FullName()
			output := m.Output().FullName()
			if input == key || output == key {
				md, found := svcDesc.MethodDescriptorByName(m.FullName())
				if found {
					return md, nil
				}
			}
		}
	}
	return nil, ErrMethodDescriptorNotFound
}
