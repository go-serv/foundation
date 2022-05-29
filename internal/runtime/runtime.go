package runtime

import (
	"errors"
	"fmt"
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
)

var (
	ErrMethodDescriptorNotFound = errors.New("")
	ErrDescriptorNotFound       = errors.New("")
)

type registryKey string
type registry map[registryKey]interface{}

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
	ref          i.ReflectInterface
	localService registry
	netServices  registry
	localClients registry
	netClients   registry
}

func (r *runtime) Reflection() i.ReflectInterface {
	return r.ref
}

func (r *runtime) ErrorWrapper(e error) {
	//switch _, e.
}

func (r *runtime) RegisterNetworkService(svc i.NetworkServiceInterface) {
	k := registryKey(svc.Name())
	if _, ok := r.netServices[k]; ok {
		panic(fmt.Sprintf("network service '%s' already registered", svc.Name()))
	}
	r.netServices[k] = svc
}

// RegisterLocalService registers a local service either by its client or by itself
func (r *runtime) RegisterLocalService(svc i.LocalServiceInterface) {
	k := registryKey(svc.Name())
	r.localService[k] = svc
}

func (r *runtime) RegisterNetworkClient(c i.NetworkClientInterface) {
	k := registryKey(c.ServiceName())
	if _, ok := r.netClients[k]; ok {
		panic(fmt.Sprintf("A network client for '%s' already registered", c.ServiceName()))
	}
	r.netClients[k] = c
}

func (r *runtime) RegisterLocalClient(c i.LocalClientInterface) {
	k := registryKey(c.ServiceName())
	if _, ok := r.localClients[k]; ok {
		panic(fmt.Sprintf("A local client for '%s' already registered", c.ServiceName()))
	}
	r.localClients[k] = c
}

func (r *runtime) RegisteredServices() []i.ServiceInterface {
	return genericRegistryAsSlice[i.ServiceInterface](r.netServices, r.localService)
}

func (r *runtime) NetworkServices() []i.NetworkServiceInterface {
	return genericRegistryAsSlice[i.NetworkServiceInterface](r.netServices)
}

func (r *runtime) NetworkClients() []i.NetworkClientInterface {
	return genericRegistryAsSlice[i.NetworkClientInterface](r.netClients)
}

func (r *runtime) LocalService() i.LocalServiceInterface {
	for _, v := range r.localService {
		return v.(i.LocalServiceInterface)
	}
	return nil
}

func (r *runtime) LocalClients() []i.LocalClientInterface {
	return genericRegistryAsSlice[i.LocalClientInterface](r.localClients)
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

func (r *runtime) ClientByMessage(msg proto.Message) (i.ClientInterface, error) {
	sf, err := r.Reflection().ServiceReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, v := range genericRegistryAsSlice[interface{}](r.localClients, r.netClients) {
		client := v.(i.ClientInterface)
		if client.ServiceName() == sf.Descriptor().FullName() {
			return client, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

func (r *runtime) ServiceByMessage(msg proto.Message) (i.ServiceInterface, error) {
	sf, err := r.Reflection().ServiceReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, v := range genericRegistryAsSlice[interface{}](r.localService, r.netServices) {
		svc := v.(i.ServiceInterface)
		if svc.Name() == sf.Descriptor().FullName() {
			return svc, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}
