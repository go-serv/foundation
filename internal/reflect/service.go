package reflect

import (
	i "github.com/go-serv/service/internal"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type service struct {
	desc      protoreflect.ServiceDescriptor
	methods   []i.MethodReflectInterface
	extValues extValueMap
}

func (s *service) Descriptor() protoreflect.ServiceDescriptor {
	return s.desc
}

func (s *service) Methods() []i.MethodReflectInterface {
	return s.methods
}

func (s *service) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return s.extValues.get(key)
}

func (s *service) Has(key *protoimpl.ExtensionInfo) bool {
	return s.extValues.has(key)
}
