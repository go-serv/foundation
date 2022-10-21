package reflect

import (
	"fmt"
	"github.com/go-serv/foundation/pkg/z"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoimpl"
	ref "reflect"
)

type serviceMapType map[protoreflect.FullName]*service

type reflect struct {
	protoExts []*protoimpl.ExtensionInfo
	services  serviceMapType
}

func (r *reflect) ServiceReflectionFromMessage(msg proto.Message) (z.ServiceReflectionInterface, error) {
	key := msg.ProtoReflect().Descriptor().FullName()
	for _, s := range r.services {
		for _, m := range s.Methods() {
			if m.Request().Descriptor().FullName() == key || m.Response().Descriptor().FullName() == key {
				return s, nil
			}
		}
	}
	return nil, fmt.Errorf("reflection: failed to find service descriptor for '%s'", key)
}

func (r *reflect) MethodReflectionFromMessage(msg proto.Message) (z.MethodReflectionInterface, error) {
	key := msg.ProtoReflect().Descriptor().FullName()
	for _, s := range r.services {
		if s == nil {
			continue
		}
		for _, m := range s.Methods() {
			if m.Request().Descriptor().FullName() == key || m.Response().Descriptor().FullName() == key {
				return m, nil
			}
		}
	}
	return nil, fmt.Errorf("reflection: failed to find method descriptor for '%s'", key)
}

// A helper function to get a service descriptor by its name
func serviceDescriptorByName(name protoreflect.FullName) (protoreflect.ServiceDescriptor, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(name)
	if err != nil {
		return nil, fmt.Errorf("reflection: %v", err)
	}
	sd, ok := desc.(protoreflect.ServiceDescriptor)
	if !ok {
		got, wanted := ref.TypeOf(sd).String(), "protoreflect.ServiceDescriptor"
		return nil, fmt.Errorf("reflection: got %s, wanted %s", got, wanted)
	}
	return sd, nil
}

func (r *reflect) AddProtoExtension(ext *protoimpl.ExtensionInfo) {
	r.protoExts = append(r.protoExts, ext)
}

func (r *reflect) Populate() error {
	var (
		err error
		sd  protoreflect.ServiceDescriptor
	)
	for name, s := range r.services {
		// Do not populate the same service again
		if s != nil {
			continue
		}
		s = new(service)
		sd, err = serviceDescriptorByName(name)
		if err != nil {
			return err
		}
		s.desc = sd
		s.extValues = r.createExtensionValueMap(sd.Options())
		r.services[name] = s
		// svc methods
		l1 := sd.Methods().Len()
		for ii := 0; ii < l1; ii++ {
			md := sd.Methods().Get(ii)
			s.methods = append(s.methods, r.newMethod(md))
		}
	}
	return nil
}

func (r *reflect) AddService(name string) {
	r.services[protoreflect.FullName(name)] = nil
}

func (r *reflect) ServiceReflectionFromName(name protoreflect.FullName) (service z.ServiceReflectionInterface, has bool) {
	service, has = r.services[name]
	return
}

type extValueMap map[*protoimpl.ExtensionInfo]interface{}

func (r *reflect) createExtensionValueMap(msg proto.Message) extValueMap {
	m := make(extValueMap, 0)
	for _, ext := range r.protoExts {
		if !proto.HasExtension(msg, ext) {
			continue
		}
		m[ext] = proto.GetExtension(msg, ext)
	}
	return m
}

func (m extValueMap) get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	for k, v := range m {
		if k == key {
			return v, true
		}
	}
	return nil, false
}

func (m extValueMap) has(key *protoimpl.ExtensionInfo) bool {
	for k, _ := range m {
		if k == key {
			return true
		}
	}
	return false
}

func (m extValueMap) bool(key *protoimpl.ExtensionInfo) bool {
	v, has := m.get(key)
	if !has {
		return false
	}
	if _, isBool := v.(bool); !isBool {
		panic(fmt.Sprintf("proto extension %s has no boolean type", key.TypeDescriptor().FullName()))
	}
	return v.(bool)
}
