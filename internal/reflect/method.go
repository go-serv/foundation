package reflect

import (
	"errors"
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"strings"
)

var (
	ErrInputOutputDescriptorNotFound = errors.New("reflection: input/output message descriptor not found")
)

func (r *reflect) newMethod(desc protoreflect.MethodDescriptor) *method {
	m := new(method)
	m.desc = desc
	m.extValues = r.createExtensionValueMap(desc.Options())
	m.req = r.newMessage(desc.Input())
	m.res = r.newMessage(desc.Output())
	return m
}

type method struct {
	desc      protoreflect.MethodDescriptor
	req       z.MessageReflectionInterface
	res       z.MessageReflectionInterface
	extValues extValueMap
}

func (m *method) Descriptor() protoreflect.MethodDescriptor {
	return m.desc
}

func (m *method) SlashFullName() string {
	name := string(m.desc.FullName())
	lastIdx := strings.LastIndex(name, ".")
	runes := []rune(name)
	runes[lastIdx] = '/'
	return "/" + string(runes)
}

func (m *method) Request() z.MessageReflectionInterface {
	return m.req
}

func (m *method) Response() z.MessageReflectionInterface {
	return m.res
}

func (m *method) IsRequest(msg proto.Message) bool {
	key := msg.ProtoReflect().Descriptor().FullName()
	switch {
	case m.req.Descriptor().FullName() == key:
		return true
	default:
		return false
	}
}

func (m *method) IsValidMessage(msg proto.Message) bool {
	return (m.IsRequest(msg) || m.IsResponse(msg))
}

func (m *method) IsResponse(msg proto.Message) bool {
	key := msg.ProtoReflect().Descriptor().FullName()
	switch {
	case m.res.Descriptor().FullName() == key:
		return true
	default:
		return false
	}
}

func (m *method) FromMessage(msg proto.Message) z.MessageReflectionInterface {
	if m.IsRequest(msg) {
		m.req.WithValue(msg)
		return m.req
	} else {
		m.res.WithValue(msg)
		return m.res
	}
}

func (m *method) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return m.extValues.get(key)
}

func (m *method) Has(key *protoimpl.ExtensionInfo) bool {
	return m.extValues.has(key)
}
