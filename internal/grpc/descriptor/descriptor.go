package descriptor

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type (
	protoExtMap map[*protoimpl.ExtensionInfo]*protoExtValue
	methodMap   map[protoreflect.FullName][]*methodDescriptor
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

type methodDescriptor struct {
	protoreflect.MethodDescriptor
	protoExts protoExtMap
}

func (m *methodDescriptor) Interface() protoreflect.MethodDescriptor {
	return m.MethodDescriptor
}

func (m *methodDescriptor) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return m.protoExts.get(key)
}

//
// Service descriptor
//

// Name returns the fully-qualified name of the service
func (s *serviceDescriptor) Name() string {
	return string(s.FullName())
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

func (r *serviceDescriptor) MethodDescriptorByName(key protoreflect.FullName) (i.MethodDescriptorInterface, bool) {
	if _, has := r.methods[key]; !has {
		return nil, false
	}
	for _, desc := range r.methods[key] {
		if desc.MethodDescriptor.FullName() == key {
			return desc, true
		}
	}
	return nil, false
}

func (r *serviceDescriptor) Descriptor() protoreflect.ServiceDescriptor {
	return r.ServiceDescriptor
}

func (r *serviceDescriptor) Populate() {
	m := r.ServiceDescriptor.Options()
	r.protoExts.populate(m.(proto.Message))
	// Methods
	methods := r.ServiceDescriptor.Methods()
	for i := 0; i < methods.Len(); i++ {
		descriptor := methods.Get(i)
		mDesc := NewMethodDescriptor(descriptor, r.methodProtoExts)
		key := mDesc.MethodDescriptor.FullName()
		if _, has := r.methods[key]; !has {
			r.methods[key] = make([]*methodDescriptor, methods.Len())
		}
		r.methods[key][i] = mDesc
	}
}
