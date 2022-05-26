package descriptor

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
)

func NewMessageDescriptor(msg proto.Message) *messageDescriptor {
	d := new(messageDescriptor)
	d.Message = msg
	return d
}

func NewMethodDescriptor(desc protoreflect.MethodDescriptor, protoExts protoExtMap) *methodDescriptor {
	r := new(methodDescriptor)
	r.MethodDescriptor = desc
	r.protoExts = protoExts
	r.protoExts.populate(desc.Options().(proto.Message))
	return r
}

func NewServiceDescriptor(svcName string) *serviceDescriptor {
	r := new(serviceDescriptor)
	r.ServiceDescriptor = newServiceDescriptor(svcName)
	r.protoExts = make(protoExtMap)
	r.methodProtoExts = make(protoExtMap)
	r.methods = make(methodMap)
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
		got, wanted := reflect.TypeOf(sd).String(), "protoreflect.ServiceDescriptorInterface"
		panic(fmt.Errorf("protobuf: got %s, wanted %s", got, wanted))
	}
	return sd
}
