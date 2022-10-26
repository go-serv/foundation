package runtime

import (
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/ancillary/memoize"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	resolversMapTyp map[any]memoize.MemoizerInterface
	eventsMapTyp    map[any][]z.EventHandlerFn
)

type runtime struct {
	platform  z.PlatformInterface
	services  map[string]z.ServiceInterface
	clients   map[string]z.ClientInterface
	resolvers resolversMapTyp
	eventsMap eventsMapTyp
}

func Runtime() *runtime {
	return rt
}

func (r *runtime) RegisterService(svc z.ServiceInterface) {
	r.services[svc.Name()] = svc
}

func (r *runtime) Services() []z.ServiceInterface {
	out := make([]z.ServiceInterface, len(r.services))
	i := 0
	for _, svc := range r.services {
		out[i] = svc
		i++
	}
	return out
}

func (r *runtime) RegisterEventHandler(eventType any, h z.EventHandlerFn) {
	if _, has := r.eventsMap[eventType]; !has {
		r.eventsMap[eventType] = make([]z.EventHandlerFn, 0)
	}
	r.eventsMap[eventType] = append(r.eventsMap[eventType], h)
}

func (r *runtime) TriggerEvent(event any, extraArgs ...any) {
	if handlers, has := r.eventsMap[event]; has {
		for i := 0; i < len(handlers); i++ {
			if stop := handlers[i](extraArgs...); stop {
				return
			}
		}
	}
}

func (r *runtime) RegisterResolver(key any, resolver memoize.MemoizerInterface) {
	r.resolvers[key] = resolver
}

func (r *runtime) Resolve(key any, args ...any) (v any, err error) {
	if resolver, ok := r.resolvers[key]; ok {
		v, err = resolver.Run(args...)
	} else {
		return nil, nil
	}
	return
}

func (r *runtime) IsRequestMessage(msg proto.Message) (bool, error) {
	m, err := service.Reflection().MethodReflectionFromMessage(msg)
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
	sf, err := service.Reflection().ServiceReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, client := range r.clients {
		name := protoreflect.FullName(client.ServiceName())
		if name == sf.Descriptor().FullName() {
			return client, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

func (r *runtime) ServiceByMessage(msg proto.Message) (z.ServiceInterface, error) {
	sf, err := service.Reflection().ServiceReflectionFromMessage(msg)
	if err != nil {
		return nil, err
	}
	for _, svc := range r.services {
		name := protoreflect.FullName(svc.Name())
		if name == sf.Descriptor().FullName() {
			return svc, nil
		}
	}
	return nil, ErrMethodDescriptorNotFound
}

//
func (r *runtime) Platform() z.PlatformInterface {
	return r.platform
}

func (r *runtime) RegisterClient(c z.NetworkClientInterface) {
	k := c.ServiceName()
	if _, ok := r.clients[k]; ok {
		// TODO: issue a warning?
		// panic(fmt.Sprintf("a client for '%s' is already registered", c.ServiceName()))
	}
	r.clients[k] = c
}
