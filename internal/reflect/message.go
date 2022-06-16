package reflect

import (
	"github.com/go-serv/service/pkg/z"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func (r *reflect) newMessage(desc protoreflect.MessageDescriptor) *message {
	m := new(message)
	m.desc = desc
	m.extValues = r.createExtensionValueMap(desc.Options())
	l1 := desc.Fields().Len()
	for k := 0; k < l1; k++ {
		fd := desc.Fields().Get(k)
		m.fields = append(m.fields, r.newField(fd))
	}
	return m
}

type message struct {
	desc      protoreflect.MessageDescriptor
	fields    []z.FieldReflectInterface
	extValues extValueMap
	value     proto.Message
}

func (m *message) Descriptor() protoreflect.MessageDescriptor {
	return m.desc
}

func (m *message) Fields() []z.FieldReflectInterface {
	return m.fields
}

func (m *message) Value() proto.Message {
	return m.value
}

func (m *message) WithValue(v proto.Message) {
	m.value = v
}

func (m *message) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return m.extValues.get(key)
}

func (m *message) Has(key *protoimpl.ExtensionInfo) bool {
	return m.extValues.has(key)
}

func (m *message) Bool(key *protoimpl.ExtensionInfo) bool {
	return m.extValues.bool(key)
}
