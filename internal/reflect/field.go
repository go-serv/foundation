package reflect

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type field struct {
	desc      protoreflect.FieldDescriptor
	extValues extValueMap
}

func (r *reflect) newField(desc protoreflect.FieldDescriptor) *field {
	f := new(field)
	f.desc = desc
	return f
}

func (f *field) Descriptor() protoreflect.FieldDescriptor {
	return f.desc
}

func (f *field) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return f.extValues.get(key)
}

func (f *field) Has(key *protoimpl.ExtensionInfo) bool {
	return f.extValues.has(key)
}
