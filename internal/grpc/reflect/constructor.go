package reflect

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
	"strings"
)

func NewMethodReflection(methodName string) *methodReflection {
	r := new(methodReflection)
	r.name = methodName
	parts := strings.Split(methodName, ".")
	r.shortName = parts[len(parts)-1]
	return r
}

func NewServiceReflection(svcName string) {
	sd := newServiceDescriptor(svcName)

	ml := sd.Methods().Len()
	for i := 0; i < ml; i++ {
		m := sd.Methods().Get(i)
		fullName := string(m.FullName())
		parts := strings.Split(fullName, ".")
		shortName := parts[len(parts)-1]
		desc := m.Options()
		if _, has := methodOpts[shortName]; !has {
			continue
		}
		for ext, opt := range methodOpts[shortName] {
			m := desc.(proto.Message)
			if !proto.HasExtension(m, ext) {
				continue
			}
			v := proto.GetExtension(m, ext)
			opt.setValue(v)
		}
	}
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

func NewMethodLevelOptions(methods []string) methodOptionsMap {
	opts := make(methodOptionsMap, 0)
	for _, name := range methods {
		opts[name] = make(optionsMap, 0)
	}
	return opts
}

func NewSvcLevelOptions() serviceOptions {
	m := make(serviceOptions)
	return m
}
