package reflect

import (
	"errors"
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
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
	req       i.MessageReflectInterface
	res       i.MessageReflectInterface
	extValues extValueMap
}

func (m *method) Descriptor() protoreflect.MethodDescriptor {
	return m.desc
}

func (m *method) Request() i.MessageReflectInterface {
	return m.req
}

func (m *method) Response() i.MessageReflectInterface {
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

func (m *method) FromMessage(msg proto.Message) i.MessageReflectInterface {
	if m.IsRequest(msg) {
		return m.req
	} else {
		return m.res
	}
}

func (m *method) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return m.extValues.get(key)
}

func (m *method) Has(key *protoimpl.ExtensionInfo) bool {
	return m.extValues.has(key)
}
