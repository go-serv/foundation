package msg

import (
	"github.com/go-serv/service/internal/runtime"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
)

type Reflection struct {
	method z.MethodReflectionInterface
	msg    z.MessageReflectionInterface
}

func (ref *Reflection) Populate(msg proto.Message) (err error) {
	reflect := runtime.Runtime().Reflection()
	if ref.method, err = reflect.MethodReflectionFromMessage(msg); err != nil {
		return
	}
	ref.msg = ref.method.FromMessage(msg)
	return
}

func (ref *Reflection) MethodReflection() z.MethodReflectionInterface {
	return ref.method
}

func (ref *Reflection) MessageReflection() z.MessageReflectionInterface {
	return ref.msg
}
