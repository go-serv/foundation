package runtime

import (
	"errors"
	"fmt"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMethodDescriptorNotFound = errors.New("")
	ErrDescriptorNotFound       = errors.New("")
)

type (
	registryKey string
	registry    map[registryKey]interface{}
)

type (
	eventsMapTyp   map[interface{}][]eventHandlerFn
	eventHandlerFn func(...interface{}) bool
)

//type registryConstraints interface {
//	i.LocalServiceInterface | i.NetworkServiceInterface | i.LocalClientInterface | i.NetworkClientInterface
//}

func genericRegistryAsSlice[T any](in ...registry) []T {
	out := make([]T, 0)
	for _, reg := range in {
		for _, item := range reg {
			out = append(out, item.(T))
		}
	}
	return out
}

type runtime struct {
	ref          z.ReflectInterface
	localService registry
	netServices  registry
	localClients registry
	netClients   registry
	eventsMap    eventsMapTyp
}

func (r *runtime) Reflection() z.ReflectInterface {
	return r.ref
}

func (r *runtime) RegisterNetworkService(svc z.NetworkServiceInterface) {
	k := registryKey(svc.Name())
	if _, ok := r.netServices[k]; ok {
		panic(fmt.Sprintf("network service '%s' already registered", svc.Name()))
	}
	r.netServices[k] = svc
}

// RegisterLocalService registers a local service either by its client or by itself
func (r *runtime) RegisterLocalService(svc z.LocalServiceInterface) {
	k := registryKey(svc.Name())
	r.localService[k] = svc
}

func (r *runtime) RegisterNetworkClient(c z.NetworkClientInterface) {
	k := registryKey(c.ServiceName())
	if _, ok := r.netClients[k]; ok {
		panic(fmt.Sprintf("A network client for '%s' already registered", c.ServiceName()))
	}
	r.netClients[k] = c
}

func (r *runtime) RegisterLocalClient(c z.LocalClientInterface) {
	k := registryKey(c.ServiceName())
	if _, ok := r.localClients[k]; ok {
		//panic(fmt.Sprintf("A local client for '%s' already registered", c.ServiceName()))
	}
	r.localClients[k] = c
}

func (r *runtime) RegisteredServices() []z.ServiceInterface {
	return genericRegistryAsSlice[z.ServiceInterface](r.netServices, r.localService)
}

func (r *runtime) NetworkServices() []z.NetworkServiceInterface {
	return genericRegistryAsSlice[z.NetworkServiceInterface](r.netServices)
}

func (r *runtime) NetworkClients() []z.NetworkClientInterface {
	return genericRegistryAsSlice[z.NetworkClientInterface](r.netClients)
}

func (r *runtime) LocalService() z.LocalServiceInterface {
	for _, v := range r.localService {
		return v.(z.LocalServiceInterface)
	}
	return nil
}

func (r *runtime) LocalClients() []z.LocalClientInterface {
	return genericRegistryAsSlice[z.LocalClientInterface](r.localClients)
}

func Runtime() *runtime {
	return rt
}

func (r *runtime) IsRequestMessage(msg proto.Message) (bool, error) {
	m, err := r.Reflection().MethodReflectionFromMessage(msg)
	if err != nil {
		return false, err
	}
	return m.IsRequest(msg), nil
}

func (r *runtime) IsResponseMessage(msg proto.Message) (bool, error) {
	ok, err := r.IsRequestMessage(msg)
	return !ok, err
}

func (r *runtime) ClientByMessage(msg proto.Message) (z.ClientInterface, error) {
	sf, err := r.Reflection().ServiceReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, v := range genericRegistryAsSlice[interface{}](r.localClients, r.netClients) {
		client := v.(z.ClientInterface)
		if client.ServiceName() == sf.Descriptor().FullName() {
			return client, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

func (r *runtime) ServiceByMessage(msg proto.Message) (z.ServiceInterface, error) {
	sf, err := r.Reflection().ServiceReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, v := range genericRegistryAsSlice[interface{}](r.localService, r.netServices) {
		svc := v.(z.ServiceInterface)
		if svc.Name() == sf.Descriptor().FullName() {
			return svc, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

func (r *runtime) RegisterEventHandler(event interface{}, h eventHandlerFn) {
	if _, has := r.eventsMap[event]; !has {
		r.eventsMap[event] = make([]eventHandlerFn, 0)
	}
	r.eventsMap[event] = append(r.eventsMap[event], h)
}

func (r *runtime) TriggerEvent(event interface{}, extra ...interface{}) {
	handlers, has := r.eventsMap[event]
	if !has {
		return
	}
	for i := 0; i < len(handlers); i++ {
		stop := handlers[i](extra...)
		if stop {
			return
		}
	}
}
