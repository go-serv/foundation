package descriptor

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
)

func NewMethodReflection(desc protoreflect.MethodDescriptor, protoExts protoExtMap) *methodReflection {
	r := new(methodReflection)
	r.name = string(desc.FullName())
	r.shortName = string(desc.Name())
	return r
}

func NewServiceReflection(svcName string) *serviceDescriptor {
	r := new(serviceDescriptor)
	r.methods = make(methodMap)
	r.ServiceDescriptor = newServiceDescriptor(svcName)
	return r
}

func newServiceDescriptor(svcFullName string) protoreflect.ServiceDescriptor {
	fullName := protoreflect.FullName(svcFullName)
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(fullName)
	if err != nil {
		panic(err)
	}
	sd, ok := desc.(protoreflect.ServiceDescriptor)
	if !ok {
		got, wanted := reflect.TypeOf(sd).String(), "protoreflect.ServiceDescriptor"
		panic(fmt.Errorf("protobuf: got %s, wanted %s", got, wanted))
	}
	return sd
}
