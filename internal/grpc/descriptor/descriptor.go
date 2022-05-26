package descriptor

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type (
	protoExtMap map[*protoimpl.ExtensionInfo]*protoExtValue
	methodMap   map[string][]*methodDescriptor
)

type protoExtValue struct {
	isSet bool
	val   interface{}
}

func (m protoExtMap) populate(msg proto.Message) {
	for ext, i := range m {
		if !proto.HasExtension(msg, ext) {
			continue
		}
		v := proto.GetExtension(msg, ext)
		i.isSet = true
		i.val = v
	}
}

func (m protoExtMap) get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	for k, v := range m {
		if k == key && v.isSet {
			return v.val, true
		}
	}
	return nil, false
}

type serviceDescriptor struct {
	protoreflect.ServiceDescriptor
	protoExts       protoExtMap
	methodProtoExts protoExtMap
	methods         methodMap
}

type messageDescriptor struct {
	proto.Message
}

func (m *messageDescriptor) MethodName() string {
	return ""
}

type methodDescriptor struct {
	protoreflect.MethodDescriptor
	protoExts protoExtMap
}

func (m *methodDescriptor) Name() string {
	return string(m.MethodDescriptor.FullName())
}

func (m *methodDescriptor) ShortName() string {
	return string(m.MethodDescriptor.Name())
}

func (m *methodDescriptor) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return m.protoExts.get(key)
}

func (s *serviceDescriptor) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return s.protoExts.get(key)
}

func (r *serviceDescriptor) AddServiceProtoExt(ext *protoimpl.ExtensionInfo) {
	v := &protoExtValue{}
	r.protoExts[ext] = v
}

func (r *serviceDescriptor) AddMethodProtoExt(ext *protoimpl.ExtensionInfo) {
	v := &protoExtValue{}
	r.methodProtoExts[ext] = v
}

func (r *serviceDescriptor) FindMethodDescriptorByName(fullName string) *methodDescriptor {
	if _, has := r.methods[fullName]; !has {
		return nil
	}
	for _, desc := range r.methods[fullName] {
		if desc.Name() == fullName {
			return desc
		}
	}
	return nil
}

func (r *serviceDescriptor) Populate() {
	m := r.ServiceDescriptor.Options()
	r.protoExts.populate(m.(proto.Message))
	// Methods
	methods := r.ServiceDescriptor.Methods()
	for i := 0; i < methods.Len(); i++ {
		descriptor := methods.Get(i)
		mDesc := NewMethodDescriptor(descriptor, r.methodProtoExts)
		key := mDesc.Name()
		if _, has := r.methods[key]; !has {
			r.methods[key] = make([]*methodDescriptor, methods.Len())
		}
		r.methods[key][i] = mDesc
	}
}
