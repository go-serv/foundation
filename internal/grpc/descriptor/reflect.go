package descriptor

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
)

type (
	protoExtMap map[*protoimpl.ExtensionInfo]*protoExtValue
	methodMap   map[string][]*methodReflection
)

type serviceDescriptor struct {
	protoreflect.ServiceDescriptor
	protoExts       protoExtMap
	methodProtoExts protoExtMap
	methods         methodMap
}

type methodReflection struct {
	name      string
	shortName string
	protoExts protoExtMap
}

func (r *serviceDescriptor) AddServiceProtoExt(ext *protoimpl.ExtensionInfo, defaultValue interface{}) {
	v := &protoExtValue{}
	v.setValue(defaultValue)
	r.protoExts[ext] = v
}

func (r *serviceDescriptor) AddMethodProtoExt(ext *protoimpl.ExtensionInfo, defaultValue interface{}) {
	v := &protoExtValue{}
	v.setValue(defaultValue)
	r.methodProtoExts[ext] = v
}

func retrieveExtValues(m proto.Message, extMap protoExtMap) {
	for ext, val := range extMap {
		if !proto.HasExtension(m, ext) {
			continue
		}
		v := proto.GetExtension(m, ext)
		val.isSet = true
		val.setValue(v)
	}
}

func (r *serviceDescriptor) Populate() {
	m := r.ServiceDescriptor.Options()
	retrieveExtValues(m.(proto.Message), r.protoExts)
	// Methods
	// Populate methods map
	methods := r.ServiceDescriptor.Methods()
	for i := 0; i < methods.Len(); i++ {
		descriptor := methods.Get(i)
		mReflect := NewMethodReflection(descriptor, r.methodProtoExts)
		retrieveExtValues(descriptor.Options(), mReflect.protoExts)
	}
}

type protoExtValue struct {
	isSet bool
	val   interface{}
}

func (o *protoExtValue) setValue(v interface{}) {
	reflect.ValueOf(o.val).Elem().Set(reflect.ValueOf(v))
}

func (o *protoExtValue) WasSet() bool {
	return o.isSet
}

func (o *protoExtValue) Value() interface{} {
	return o.val
}

func (o protoExtMap) OverrideIfNotSet(v interface{}, targetExt *protoimpl.ExtensionInfo) {
	for ext, val := range o {
		if ext == targetExt && !val.isSet {
			val.setValue(v)
		}
	}
}
