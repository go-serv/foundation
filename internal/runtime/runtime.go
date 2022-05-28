package runtime

import (
	"errors"
	"fmt"
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
	localService registry
	netServices  registry
	localClients registry
	netClients   registry
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

func (r *runtime) RegisterLocalService(svc i.LocalServiceInterface) {
	k := registryKey(svc.Name())
	if _, ok := r.localClients[k]; ok {
		panic(fmt.Sprintf("Only one local service is allowed per application, '%s' already registered", svc.Name()))
	}
	r.localClients[k] = svc
}

func (r *runtime) RegisterNetworkClient(c i.NetworkClientInterface) {
	k := registryKey(c.ServiceName())
	if _, ok := r.netClients[k]; ok {
		panic(fmt.Sprintf("A network client for '%s' already registered", c.ServiceName()))
	}
	r.netClients[k] = c
}

func (r *runtime) RegisterLocalClient(svcName protoreflect.FullName, c i.LocalClientInterface) {
	k := registryKey(svcName)
	if _, ok := r.localClients[k]; ok {
		panic(fmt.Sprintf("A local client for '%s' already registered", svcName))
	}
	r.localClients[k] = c
}

func (r *runtime) NetworkServices() []i.NetworkServiceInterface {
	return genericRegistryAsSlice[i.NetworkServiceInterface](r.netServices)
}

func (r *runtime) NetworkClients() []i.NetworkClientInterface {
	return genericRegistryAsSlice[i.NetworkClientInterface](r.netClients)
}

func (r *runtime) LocalService() i.LocalServiceInterface {
	//s := genericRegistryAsSlice[i.LocalClientInterface](r.localService)
	//return s[0]
	return nil
}

func (r *runtime) LocalClients() []i.LocalClientInterface {
	return genericRegistryAsSlice[i.LocalClientInterface](r.localClients)
}

func Runtime() *runtime {
	return rt
}

func (r *runtime) IsRequestMessage(msg proto.Message) (bool, error) {
	m, err := r.MethodDescriptorByMessage(msg)
	if err != nil {
		return false, err
	}
	return m.Interface().Input().FullName() == msg.ProtoReflect().Descriptor().FullName(), nil
}

func (r *runtime) IsResponseMessage(msg proto.Message) (bool, error) {
	ok, err := r.IsRequestMessage(msg)
	return !ok, err
}

func (r *runtime) ServiceDescriptorByMessage(msg proto.Message) (i.ServiceDescriptorInterface, error) {
	key := msg.ProtoReflect().Descriptor().FullName()
	//
	items := genericRegistryAsSlice[interface{}](r.netServices, r.localService)
	for _, svc := range items {
		svcDesc := svc.(i.ServiceInterface).Service_Descriptor()
		methods := svcDesc.Descriptor().Methods()
		l1 := methods.Len()
		for ii := 0; ii < l1; ii++ {
			m := methods.Get(ii)
			input := m.Input().FullName()
			output := m.Output().FullName()
			if input == key || output == key {
				return svcDesc, nil
			}
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

func (r *runtime) MethodDescriptorByMessage(msg proto.Message) (i.MethodDescriptorInterface, error) {
	key := msg.ProtoReflect().Descriptor().FullName()
	//
	items := genericRegistryAsSlice[interface{}](r.netServices, r.localService)
	for _, svc := range items {
		svcDesc := svc.(i.ServiceInterface).Service_Descriptor()
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

func (r *runtime) ClientByMessage(msg proto.Message) (i.ClientInterface, error) {
	svcDesc, err := r.ServiceDescriptorByMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, v := range genericRegistryAsSlice[interface{}](r.localClients, r.netClients) {
		client := v.(i.ClientInterface)
		if client.ServiceName() == svcDesc.Descriptor().FullName() {
			return client, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

func (r *runtime) ServiceByMessage(msg proto.Message) (i.ServiceInterface, error) {
	svcDesc, err := r.ServiceDescriptorByMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, v := range genericRegistryAsSlice[interface{}](r.localService, r.netServices) {
		svc := v.(i.ServiceInterface)
		if svc.Name() == svcDesc.Descriptor().FullName() {
			return svc, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}
