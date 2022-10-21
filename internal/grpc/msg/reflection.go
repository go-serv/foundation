package msg

import (
	"github.com/go-serv/foundation/internal/service"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/protobuf/proto"
)

type Reflection struct {
	service z.ServiceReflectionInterface
	method  z.MethodReflectionInterface
	msg     z.MessageReflectionInterface
}

func (ref *Reflection) Populate(msg proto.Message) (err error) {
	reflect := service.Reflection()
	if ref.method, err = reflect.MethodReflectionFromMessage(msg); err != nil {
		return
	}

	ref.msg = ref.method.FromMessage(msg)

	if ref.service, err = reflect.ServiceReflectionFromMessage(msg); err != nil {
		return
	}
	return
}

func (ref *Reflection) ServiceReflection() z.ServiceReflectionInterface {
	return ref.service
}

func (ref *Reflection) MethodReflection() z.MethodReflectionInterface {
	return ref.method
}

func (ref *Reflection) MessageReflection() z.MessageReflectionInterface {
	return ref.msg
}
